v_cc_library(
  NAME config
  SRCS
    configuration.cc
    base_property.cc
    rjson_serialization.cc
  DEPS
    v::json
    v::model
   boost_filesystem
)
add_subdirectory(tests)
