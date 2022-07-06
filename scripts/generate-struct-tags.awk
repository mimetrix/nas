#!/bin/awk -f

# This script insert struct tags inside the type <name type> struct.

# match type (first pattern), print, activate process and jump next line.
# TODO missing to handle when the struct is empty, e.g. type Foo struct {}
/type(.*)+struct(.*){/{
  activateProcess = 1;
  started=0
}
# match } (second pattern), deactivate process.
/}/{
  print $0
  activateProcess = 0
  started=0
  next;
}
{
  # if is not started, so print the type <name> struct{ and move to the next line.
  if(started == 0){
    print $0
    started=1
    next;
  }
  # if activate is on, create struct tag.
  if(activateProcess == 1){
    # if there is // is because there is a comment in the end (likely the $2 field).
    # if the first contains . (dot) is because there is a module chasing.
    # if the first contains * (asterist) is because it is a pointer.
    if ($2 ~ /\/\//){
      printf "%s", $1
    } else {
      printf "%s %s", $1, $2
    }
    if( $1 ~ /\./){
      FS="[. ]"
      $0=$0
      # TODO navarrothiago Check if there are more than one dot.
      if ($2 ~ /\*/ ){
        printf "  `json:\"%s,omitempty\"`", substr($2, 2)
      }else{
        printf "  `json:\"%s,omitempty\"`", $2
      }
    }else{
      if ($1 ~ /\*/ ){
        printf "  `json:\"%s,omitempty\"`", substr($1, 2)
      } else{
        printf "  `json:\"%s,omitempty\"`", $1
      }
    }
    print ""
    FS=" "
    # if contains comments, it will be removed.
  }else{
    print $0
  }
}
