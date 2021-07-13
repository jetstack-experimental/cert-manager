# Copyright 2020 The cert-manager Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

def helm_pkg(
    name,
    chart_name,
    helmignore,
    readme_file,
    chart_yaml,
    values_yaml,
    tpl_files = [],
    srcs = [],
    helm_cmd = "//hack/bin:helm",
    version_file = "//:version",
    **kwargs):

    pkg_tar(
        name = "%s.template_files" % name,
        package_dir = "/%s/templates" % chart_name,
        srcs = srcs,
        mode = "0644",
        visibility = ["//visibility:private"],
    )

    pkg_tar(
        name = "%s.tpl_files" % name,
        package_dir = "/%s/templates" % chart_name,
        srcs = tpl_files,
        mode = "0644",
        visibility = ["//visibility:private"],
    )

    pkg_tar(
        name = "%s.chart_files" % name,
        package_dir = "/%s" % chart_name,
        files = {
            helmignore: ".helmignore"
            readme_file: "README.md",
            chart_yaml: "Chart.yaml",
            values_yaml: "values.yaml",
        },
        mode = "0644",
        visibility = ["//visibility:private"],
    )

    pkg_tar(
        name = "%s.dir_tar" % name,
        extension = "tar.gz",
        deps = [
            "%s.template_files" % name,
            "%s.tpl_files" % name,
            "%s.chart_files" % name,
        ],
        visibility = ["//visibility:private"],
    )

    cmds = []
    cmds = cmds + ["tar xf $(location %s.dir_tar)" % name]
    cmds = cmds + ["version=$$(cat $(location %s))" % version_file]
    cmds = cmds + [" ".join([
        "$(location %s)" % helm_cmd,
        "package",
        "--app-version=$$version",
        "--version=$$version > /dev/null 2>&1",
        "./%s" % chart_name,
    ])]
    cmds = cmds + ["mv \"%s-$$version.tgz\" $@" % chart_name]
    native.genrule(
        name = name,
        srcs = [
            version_file,
            ":%s.dir_tar" % name,
        ],
        stamp = 1,
        outs = ["%s.tgz" % name],
        cmd = "; ".join(cmds),
        tools = [helm_cmd],
        **kwargs
    )

def helm_tmpl(
        name,
        helm_pkg,
        release_namespace,
        release_name,
        additional_api_versions = "",
        values = {},
        helm_cmd = "//hack/bin:helm",
        outs = [],
        **kwargs,
):
    tmpl_cmd = [
        "$(location %s)" % helm_cmd,
        "template",
        "--api-versions=\"%s\"" % additional_api_versions,
        "--namespace=%s" % release_namespace,
        release_name,
        "$(location %s)" % helm_pkg,
    ]

    for k, v in values.items():
        tmpl_cmd = tmpl_cmd + ["--set=\"%s=%s\"" % (k, v)]

    if len(outs) == 0:
        tmpl_cmd = tmpl_cmd + ["> $@"]
        outs = ["%s.yaml" % name]
        cmds = [" ".join(tmpl_cmd)]
    else:
        tmpl_cmd = tmpl_cmd + ["--output-dir=$(@D)"]
        cmds = [" ".join(tmpl_cmd), "mv $(@D)/%s/* $(@D)" % release_name]

    native.genrule(
        name = name,
        srcs = [helm_pkg, "//:version"],
        stamp = 1,
        outs = outs,
        cmd = "; ".join(cmds),
        tools = [helm_cmd],
        **kwargs,
    )
