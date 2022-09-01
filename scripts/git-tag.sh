git tag $(git describe | awk -F- '{print "v1.0.7-" $5 "-" $6}')
