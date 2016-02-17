package uaparser

import "testing"

import "fmt"

func TestParse(t *testing.T) {
	var expectedBrowserNames map[string][]string = GetBrowserNames()
	var expectedOperatingSystems map[string][]string = GetOSNames()
	// var expectedDeviceTypes map[string][]string = GetDeviceTypes()
	var totalTestCount = 0

	totalTestCount += _checkExepectations(t, expectedOperatingSystems, "os")
	totalTestCount += _checkExepectations(t, expectedBrowserNames, "browser")
	// totalTestCount += _checkExepectations(t, expectedDeviceTypes, "deviceType")

	fmt.Printf("Ran %d tests\n", totalTestCount)
}

func _checkExepectations(t *testing.T, expectations map[string][]string, testType string) (testCount int) {
	var uaParseResult *UAInfo
	var testResult bool
	var comparedTo string

	testCount = 0

	for expectation := range expectations {
		for key := range expectations[expectation] {
			fmt.Printf("Testing: %s \n", expectations[expectation][key])
			uaParseResult = Parse(expectations[expectation][key])

			if testType == "browser" {
				testResult, comparedTo = _checkBrowser(uaParseResult, expectation)
			} else if testType == "os" {
				testResult, comparedTo = _checkOs(uaParseResult, expectation)
			} else if testType == "deviceType" {
				testResult, comparedTo = _checkDeviceType(uaParseResult, expectation)
			}

			if !testResult {
				t.Fatalf("Expected: '%s' got '%s' on: '%s'", expectation, comparedTo, expectations[expectation][key])
			}
			testCount += 1
		}
	}

	return testCount
}

func _checkBrowser(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.Browser == nil {
		return false, ""
	}
	return (uainfo.Browser.Name == expectation), uainfo.Browser.Name
}

func _checkOs(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.OS == nil {
		return false, ""
	}
	return (uainfo.OS.Name == expectation), uainfo.OS.Name
}

func _checkDeviceType(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.DeviceType == nil {
		return false, ""
	}
	return (uainfo.DeviceType.Name == expectation), uainfo.DeviceType.Name
}

func TestParse2(t *testing.T) {
	result := Parse("Mozilla/5.0 (Windows NT 6.2; Win64; x64; Trident/7.0; rv:11.0) like Gecko")

	if result.Browser.Name != "IE" ||
		result.Browser.Version != "11.0" ||
		result.OS.Name != "Windows" {
		t.Error("wrong result")
	}
}

func TestParse3(t *testing.T) {
	result := Parse("Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 920) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537")

	if result.Browser.Name != "IE" ||
		result.Browser.Version != "11.0" ||
		result.Device.Name != "iPhone" ||
		result.OS.Name != "Windows Phone OS" {
		t.Error("wrong result")
	}
}

func TestParse4(t *testing.T) {
	result := Parse("Mozilla/5.0 (PLAYSTATION 3 4.78) AppleWebKit/531.22.8 (KHTML, like Gecko)")

	if result.Browser.Name != "PlayStation" ||
		result.OS.Name != "PlayStation OS" {
		t.Error("wrong result")
	}
}
