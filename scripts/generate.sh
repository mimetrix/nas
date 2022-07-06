#!/usr/bin/env bash
main() {

  set -o errexit
  set -o pipefail
  set -o nounset

  set -x
  local -r __dirname="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  local -r __filename="${__dirname}/$(basename "${BASH_SOURCE[0]}")"

  target="${__dirname}"/../$1

  for item in $(ls "${target}"/*.go | awk '{print $0}' | grep -v _test); do
    # skip some edges case.
    if [ "${item}" == "coverage.out" ] || [ "${item}" == "NAS_Plain5GSNASMessage.go" ]; then
      echo "Skipping "${item}""
      continue
    fi
    echo "Generating struct tag for ${item}..."
    tmp_file="${item}".tmp
    "${__dirname}"/generate-struct-tags.awk "${item}" > "${tmp_file}"
    if [ "$?" != 0 ]; then
      rm "${tmp_file}"
      echo "Error when generating struct tag for ${item}"
      exit 1
    fi
    mv "${tmp_file}" "${item}"
    gofmt -w "${item}"
  done
  exit 0
}
main "$@"
