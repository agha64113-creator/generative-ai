# How to Contribute

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Notebook Template

If you're creating a Jupyter Notebook, use [`notebook_template.ipynb`](notebook_template.ipynb) as a template.

## Code Quality Checks

All notebooks in this project are checked for formatting and style, to ensure a
consistent experience. To test notebooks prior to submitting a pull request,
you can follow these steps.

From a command-line terminal (e.g. from Vertex AI Workbench or locally),
run this code block to format your code.
If the fixes can't be performed automatically,
then you will need to manually address them before submitting your PR.

```shell
python3 -m pip install --upgrade nox
nox -s format
```

## Code Reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows [Google's Open Source Community Guidelines](https://opensource.google/conduct/).

## Contributor Guide

If you are new to contributing to open source, you can find helpful information in this contributor guide.

You may follow these steps to contribute:

1. **Fork the official repository.** This will create a copy of the official repository in your own account.
2. **Sync the branches.** This will ensure that your copy of the repository is up-to-date with the latest changes from the official repository.
3. **Work on your forked repository's feature branch.** This is where you will make your changes to the code.
4. **Commit your updates on your forked repository's feature branch.** This will save your changes to your copy of the repository.
5. **Submit a pull request to the official repository's main branch.** This will request that your changes be merged into the official repository.
6. **Resolve any linting errors.** This will ensure that your changes are formatted correctly.

### Step-by-Step Git & GitHub Workflow

The following commands walk you through the complete workflow from fork to pull request.

#### 1. Fork the repository

On GitHub, click the **Fork** button at the top-right of this page:
<https://github.com/GoogleCloudPlatform/generative-ai>

This creates your own copy at `https://github.com/<your-username>/generative-ai`.

#### 2. Clone your fork locally

```shell
git clone https://github.com/<your-username>/generative-ai.git
cd generative-ai
```

#### 3. Add the upstream remote

```shell
git remote add upstream https://github.com/GoogleCloudPlatform/generative-ai.git
```

Verify the remotes are set up correctly:

```shell
git remote -v
# origin    https://github.com/<your-username>/generative-ai.git (fetch)
# origin    https://github.com/<your-username>/generative-ai.git (push)
# upstream  https://github.com/GoogleCloudPlatform/generative-ai.git (fetch)
# upstream  https://github.com/GoogleCloudPlatform/generative-ai.git (push)
```

#### 4. Create a feature branch

Always branch off the latest `main`:

```shell
git fetch upstream
git checkout -b my-feature-branch upstream/main
```

#### 5. Make your changes

Edit files, add notebooks, or fix bugs. Then stage and commit your work:

```shell
git add .
git commit -m "feat: describe your change here"
```

For multiple logical changes, use separate commits:

```shell
git add path/to/file1
git commit -m "fix: correct typo in notebook"

git add path/to/file2
git commit -m "feat: add new sample for Gemini function calling"
```

#### 6. Keep your branch up to date

If new commits have landed on upstream `main` while you were working, rebase to avoid conflicts:

```shell
git fetch upstream
git rebase upstream/main
```

#### 7. Push your branch to your fork

```shell
git push origin my-feature-branch
```

If you previously pushed and then rebased, use `--force-with-lease` to update the remote safely:

```shell
git push --force-with-lease origin my-feature-branch
```

#### 8. Open a Pull Request

1. Go to your fork on GitHub: `https://github.com/<your-username>/generative-ai`
2. Click **Compare & pull request** (GitHub will prompt you after a push).
3. Set the **base repository** to `GoogleCloudPlatform/generative-ai` and **base branch** to `main`.
4. Fill in a clear title and description explaining *what* you changed and *why*.
5. Click **Create pull request**.

#### 9. Respond to review feedback

If reviewers request changes:

```shell
# Make the requested edits, then:
git add .
git commit -m "fix: address review feedback"
git push origin my-feature-branch
```

The PR will update automatically.

#### 10. After your PR is merged

Clean up your local and remote branches:

```shell
git checkout main
git pull upstream main
git branch -d my-feature-branch
git push origin --delete my-feature-branch
```
   - For errors generated by [check-spelling](https://github.com/check-spelling/check-spelling), go to the [Job Summary](https://github.com/GoogleCloudPlatform/generative-ai/actions/workflows/spelling.yaml) to read the errors.
     - Fix any spelling errors that are found.
     - Forbidden Patterns are defined as regular expressions, you can copy/paste them into many IDEs to find the instances. [Example for Visual Studio Code](https://medium.com/@nikhilbaxi3/visual-studio-code-secrets-of-regular-expression-search-71723c2ecbd2).
     - Add false positives to [`.github/actions/spelling/allow.txt`](.github/actions/spelling/allow.txt). Be sure to check that it's actually spelled correctly!

Here are some additional things to keep in mind during the process:

- **Read the [Google's Open Source Community Guidelines](https://opensource.google/conduct/).** The contribution guidelines will provide you with more information about the project and how to contribute.
- **Test your changes.** Before you submit a pull request, make sure that your changes work as expected.
- **Be patient.** It may take some time for your pull request to be reviewed and merged.

---

## For Google Employees

Complete the following steps to register your GitHub account and be added as a contributor to this repository.

1. Register your GitHub account at [go/GitHub](http://go/github)

1. Once you have registered, go to [go/github-googlecloudplatform](http://go/github-googlecloudplatform) and request to join the GoogleCloudPlatform organization. Check the box "I need write access on a public repository".

1. You'll receive an email to your GitHub registered email to approve the request to join. Approve it.

1. Request to join this team [GoogleCloudPlatform/teams/generative-ai-samples-contributors](https://github.com/orgs/GoogleCloudPlatform/teams/generative-ai-samples-contributors/members)
