// Copyright 2022 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code originally generated by "mdtogo", but no
// longer maintained that way.
package docs

var RepoShort = `Manage package repositories.`
var RepoLong = `
The ` + "`" + `repo` + "`" + ` command group contains subcommands for managing package repositories.
`

var GetShort = `List registered repositories.`
var GetLong = `
  porchctl repo get [REPOSITORY_NAME] [flags]

Args:

  REPOSITORY_NAME:
    The name of a repository. If provided, only that specific
    repository will be shown. Defaults to showing all registered
    repositories.
`
var GetExamples = `
  # list all repositories registered in the default namespace
  $ porchctl repo get --namespace default

  # show the repository named foo in the bar namespace
  $ porchctl repo get foo --namespace bar
`

var RegShort = `Register a package repository.`
var RegLong = `
  porchctl repo reg REPOSITORY [flags]

Args:

  REPOSITORY:
    The URI for the registry. It can be either a git repository
    or an oci repository. For the latter, the URI must have the
    'oci://' prefix.

Flags:

  --branch:
    Branch within the repository where finalized packages are
    commited. The default is to use the 'main' branch.
  
  --deployment:
    Tags the repository as a deployment repository. Packages in
    a deployment repository are considered ready for deployment.
  
  --description:
    Description of the repository.
  
  --directory:
    Directory within the repository where packages are found. The
    default is the root of the repository.
  
  --name:
    Name of the repository. By default the last segment of the
    repository URL will be used as the name.
  
  --repo-basic-username:
    Username for authenticating to a repository with basic auth.
  
  --repo-basic-password:
    Password for authenticating to a repository with basic auth.
`
var RegExamples = `
  # register a new git repository with the name generated from the URI.
  $ porchctl repo register https://github.com/platkrm/demo-blueprints.git --namespace=default

  # register a new deployment repository with name foo.
  $ porchctl repo register https://github.com/platkrm/blueprints-deployment.git --name=foo --deployment --namespace=bar
`

var UnregShort = `Unregister a repository.`
var UnregLong = `
  porchctl repo unreg REPOSITORY_NAME [flags]

Args:

  REPOSITORY_NAME:
    The name of a repository.

Flags:

  --keep-auth-secret:
    Keep the Secret object with auth information referenced by the repository.
    By default, it will be deleted when the repository is unregistered.
`
var UnregExamples = `
  # unregister a repository and keep the auth secret.
  $ porchctl repo unreg registered-repository --namespace=default --keep-auth-secret
`