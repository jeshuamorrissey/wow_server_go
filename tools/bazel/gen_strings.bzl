"""Utility macro for runner `stringer` across a set of files."""

def gen_strings(name, srcs, types):
    """Generate a string representation of the types using "stringer".

    Args:
        name: The name of the rule.
        srcs: Golang source files to include.
        types: A list of types to generate strings for.
    """
    all_files = []
    for type in types:
        rule_name = "{}.{}".format(name, type)
        filename = "{}.go".format(rule_name)
        all_files.append(filename)

        native.genrule(
            name = rule_name,
            srcs = srcs,
            outs = [filename],
            cmd = "HOME=$$(pwd) $(location @org_golang_x_tools//cmd/stringer) -output $@ -trimprefix {} -type {} {}/*".format(
                type,
                type,
                native.package_name(),
            ),
            tools = ["@org_golang_x_tools//cmd/stringer"],
        )

    native.filegroup(
        name = name,
        srcs = all_files,
    )
