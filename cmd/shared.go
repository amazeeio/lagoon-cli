package cmd

import (
	"os"
	"sort"
	"strings"

	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/hashicorp/go-version"
)

// config vars
var lagoonHostname string
var lagoonPort string
var lagoonGraphQL string
var lagoonToken string
var lagoonUI string
var lagoonKibana string

// variable vars
var variableValue string
var variableName string
var variableScope string

// user vars
var userFirstName string
var userLastName string
var userEmail string
var pubKeyFile string
var nameInPubKeyFile bool
var sshKeyName string

// group vars
var groupName string
var groupRole string

var jsonPatch string
var revealValue bool
var listAllProjects bool
var noHeader bool

// These are available to all cmds and are set either by flags (-p and -e) or via `lagoon-cli/app` when entering a directory that has a valid lagoon project
var cmdProjectName string
var cmdProjectEnvironment string

var remoteID string

var notificationName string
var notificationNewName string
var notificationWebhook string
var notificationChannel string

var deployBranchName string

var outputOptions = output.Options{
	Header: false,
	CSV:    false,
	JSON:   false,
	Pretty: false,
	Debug:  false,
}

var debugEnable bool

func handleError(err error) {
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}

func returnNonEmptyString(value string) string {
	if len(value) == 0 {
		return "-"
	}
	return value
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func stripNewLines(stripString string) string {
	return strings.TrimSuffix(stripString, "\n")
}

// sort fields of a given map to iterate over returned fields from an api object
func sortFields(m map[string]string, fields *[]string) {
	for k := range m {
		*fields = append(*fields, k)
	}
	sort.Strings(*fields)
}

// insert a string into a slice
func insertString(array []string, value string, index int) []string {
	return append(array[:index], append([]string{value}, array[index:]...)...)
}

// remove a string from a slice
func removeString(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

// move a string around in a slice
func moveString(array []string, srcIndex int, dstIndex int) []string {
	value := array[srcIndex]
	return insertString(removeString(array, srcIndex), value, dstIndex)
}

// check if slice contains a specific string
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// get index of element in slice
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// remove from a slice if it exists
func removeFromSlice(array []string, element string) []string {
	if containsString(array, element) {
		idx := sliceIndex(len(array), func(i int) bool { return array[i] == element })
		array = removeString(array, idx)
	}
	return array
}

// return fields we always want to show
func alwaysShowFields(show []string, always map[int]string) []string {
	keys := make([]int, 0)
	for k := range always {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if !containsString(show, always[k]) {
			show = append(show, always[k])
		}
	}
	for _, k := range keys {
		// and to make sure they are in the first positions, not randomly through the response
		show = moveString(show, sliceIndex(len(show), func(i int) bool { return (show)[i] == always[k] }), k)
	}
	return show
}

// return if the version is greater or less than
func greaterThanOrEqualVersion(a string, b string) bool {
	aVer, err := version.NewSemver(a)
	if err != nil {
		return false
	}
	bVer, err := version.NewSemver(b)
	if err != nil {
		return false
	}
	if aVer.GreaterThanOrEqual(bVer) {
		return true
	}
	return false
}
