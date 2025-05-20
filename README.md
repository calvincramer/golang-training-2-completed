# Golang Training 2 - Intermediate, Functional Mindset
In `intermediate.go` there are exercises to complete. Try completing them without looking at the tests.

In this training there are some test cases that are commented out since they would otherwise cause compilation errors before the corresponding code in `intermediate.go` in completed. Look for the `TODO`s in `intermediate_test.go` and uncomment the test cases as you are working on the exercises.

Run the tests with `go test` or `go test -v`.

Passing implementations will be made available at a later time [here](https://github.com/calvincramer/golang-training-2-completed).

# Install Golang (Linux):
1. Download from https://go.dev/dl/
2. Run:
```sh
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz  # use the file downloaded from step 1
```
3. Make sure this path is setup correctly: `export PATH=$PATH:/usr/local/go/bin`. Add this in `.bashrc` or equivalent for your shell.
4. Start a new terminal or reload the current shell environment
5. Verify go works by running `go version`
