package main

/*
#include <stdbool.h>
struct LabelSelectorMatchesResult {
    bool matches;
    bool has_error;
    char* error_message;
};
*/
import "C"
import "k8s.io/apimachinery/pkg/labels"
import "encoding/json"

//export matches_label_selector
func matches_label_selector(labelsJSON, labelSelectorStr *C.char) C.struct_LabelSelectorMatchesResult {
   var labelsMap map[string]string
   err := json.Unmarshal([]byte(C.GoString(labelsJSON)), &labelsMap)
   if err != nil {
      return C.struct_LabelSelectorMatchesResult{C.bool(false), C.bool(true), C.CString(err.Error())}
   }

   labelSelector, err := labels.Parse(C.GoString(labelSelectorStr))
	if err != nil {
		return C.struct_LabelSelectorMatchesResult{C.bool(false), C.bool(true), C.CString(err.Error())}
	}

   return C.struct_LabelSelectorMatchesResult{C.bool(labelSelector.Matches(labels.Set(labelsMap))), C.bool(false), C.CString("")}
}


func main() {
// This is necessary for the compiler.
// You can add something that will be executed when engaging your library to the interpreter.
}
