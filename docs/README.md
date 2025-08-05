# Docs

## Commands

`tf-changelog init` will initialize your repo to be used with `tf-changelog`.
This will create a `mkdocs.yml` file, a `docs/` directory if it doesn't already
exist, and follow the default `docs/` creating files or directories that don't
exist.

`tf-changelog update --all` will update all module docs with any new releases.
It will also update the `CHANGELOG.md` with any changes it doesn't have.

`tf-changelog release` will create tags for any modules that have been updated.

## Mono Repo Layout

An example layout and modules are in the `examples/tf-modules` directory. This
is to represent a Github Repo that follows a pattern used widely with tools like
`Terragrunt`.

Inside the repo the layout looks like this:

```sh
.
├── docs
│   ├── index.md
│   └── modules
│       ├── bucket.md
│       └── vpc.md
├── mkdocs.yml
├── modules
│   ├── bucket
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   ├── README.md
│   │   └── variables.tf
│   └── vpc
│       ├── main.tf
│       ├── outputs.tf
│       ├── README.md
│       └── variables.tf
└── README.md

```

## Terraform modules docs

Each Terraform module will generate an output equivalent to running
[terraform-docs](https://github.com/terraform-docs/terraform-docs). The template
for the output will also include usage docs for the module in form of
`Terragrunt` call style. The top of the file will include a table that shows the
most recent versions of that module.

### Docs layout

The `docs/` layout in your modules repo should follow the
[Backstage Techdocs](https://backstage.io/docs/features/techdocs/creating-and-publishing#writing-and-previewing-your-documentation)
setup that looks like this:

```sh
├── docs
│   ├── index.md
│   └── modules
│       ├── bucket.md
│       └── vpc.md
├── mkdocs.yml
```

The navigation and addons are from [mkdocs](https://www.mkdocs.org/).

### Mkdocs

The default `mkdocs.yml` looks like this:

```markdown
site_name: My Terraform Modules site_description: My Terraform Modules Docs
plugins:

- techdocs-core nav:
- Getting Started: index.md
- Latest Changes: CHANGELOG.md
- Modules:
  - vpc: modules/vpc.md
  - bucket: modules/bucket.md
```

- `index.md` is the Home page for the documentation.
- `CHANGELOG.md` is a complete changelog file of all running changes - newest at
  the top.
- `modules/` contains a markdown file of the name of the module, ie : `vpc.md`

### Docs templates

The Terraform modules are generated using
[terraform-docs](https://github.com/terraform-docs/terraform-docs) as stated
before. The `index.md` file is updated each time `tf-changelog` is ran to add or
modify any new modules to the `nav`. The `CHANGELOG.md` is also generated to add
the newest module tags and changes to the file.

## Release tags

Any module with a change in it at runtime will generate a new module release
tag. These are different from the github release tags in they are prefixed by
the path of the module. For example:

If the module path is `modules/vpc` then the tag will be `vpc/v0.0.1`. If the
module path is `modules/aws/vpc` then the tag by default will be
`aws/vpc/v0.0.1`

You can modify the prefix that is trimmed in the `.tf-changelog.yaml`
configuration.

All module versions will have a semantic versioning suffix of `vX.X.X`
