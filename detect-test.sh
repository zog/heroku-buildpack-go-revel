it_is_revel_if_conf_files_exist() {
  sh -x bin/detect test
}

it_is_not_revel_if_conf_files_do_not_exist() {
  ! sh -x bin/detect .
}
