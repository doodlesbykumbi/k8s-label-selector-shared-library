require 'ffi'
require 'json'

class LabelSelectorMatchesResult < FFI::Struct 
    layout  :matches, :bool,
            :has_error, :bool, 
            :error_message, :string
end

module GoSharedLib
  extend FFI::Library
  ffi_lib './go_shared_lib.so'

  attach_function :matches_label_selector, [:string, :string], LabelSelectorMatchesResult.by_value
end

namespace_labels = {:"field.cattle.io/projectId" => "p-zvd84", :"kubernetes.io/metadata.name" => "fleet-local"}
label_selector = "field.cattle.io/projectId in (123, p-zvd84)"

matches_result = GoSharedLib.matches_label_selector(namespace_labels.to_json, label_selector)
puts "matches=#{matches_result[:matches]}, has_error=#{matches_result[:has_error]}, error_message=#{matches_result[:error_message]}"
