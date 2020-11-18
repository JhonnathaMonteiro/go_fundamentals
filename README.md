Writing tests

Writing a test just like writing a function, with a few rules
  - It needs to be in a file with name like xxx_test.go
  - The test function must start with the word Test
  - The test function takes one argument only t *testing.T

For now it's enought to kno that your t type *testing.T is your "hook" into the testing framework
so you can do things like t.Fail() when you want to fail.

We've covered some new topics:

if

If statements in Go are very much like other programming languages

##Declaring variables

We're declaring some variables with the sintax varName := value, which lets us re0use some values in our test
for readibility

  t.Errorf

  We are calling the Errorf method on our t which will print out a message and fail the test. The f
  stands for format which allows us to build a string with values inserted into the placeholder values %q.
  When you made the test fail it should be clear how it works

  You can read more about placeholder sTrings into the fmt go doc. For tests %q is very useful as it wraps
  your values in double quotes.

  We will later explore tyhe diverence between methods and functions.

  Go doc

  Another quality of life feature of Go is the documentation. You can launch the docs locally by running
  godoc -http :8000. If you go to localhost:8000/pkg you will see all the packages intealled on your system.

  The vast majority of the standard library has excellent documentation with examples. Navigating to
  http://localhost:8000/pkg/testing would be a worthwhile to see what's available to you.

  If you don't have godoc command, than maybe you are using the newer version of Go (1.14 or later)
  which is no longer including godoc. You can manually install it with go get dolang.org/x/tools/cmd/godoc

  Hello, YOU

  Now that we have a test we can iterate on our software safetly

  In the last example we wrote the test after the code had been written just so you could get an example of
  how to write a test and declare a function. From this point on we will be writing tests first.

  Our next requirement is to let us specify the recipient of the greeting.

  Let's start by capturing these requirements in a test. This is basic the test driven development and allow us
  to make sure our test is actually testing what we want. When you retrospectively write tests there is the
  risk that your thest may continue to pass even if the code doesn't work as intended.

  ==CODE==

  Now run go test, you should have a compilation error

  ==CODE==

  When using a statically typed language like Go it is important to listen to the compiler. The compiler
  understands how your code should snap together and work so you don't have to.

  In this case the compiler is telling you what you need to do to continue. We have to change our function
  Hello to accept an argument.

  Edit the Hello function to accept an argument of type string

  ==CODE==

  If you try and run your tests again your main.go will fail to compile because you're not passing an 
  argument. Send in "world" to make it pass.

  ==CODE==

  Now when you run your tests you should see something like

  ==CODE==

  We finally have a compiling program but it is not meeting our requirements according to the test.

  Let's make the test pass by using the name argument and concatenatge it with Hello,

  ==CODE==

  When you run the tests they should now pass. Normally as part of the TDD cycle we should now refactor.

  ##A note on source control

  At this point, if you are using source control (wich you should!) I would commit the code as it if. We
  have working software backed by a test.

  I wouldn't push to master though, because I plan to refactor next. It is nice to commit at this point in case
  you somehow get into a mess with refactoring - you can always go back to the working version.

  There's not a lot to refactor here, but we can introduce another language feature, constants.

## Constants
  
  Constants are defined like so

  ```const englishHelloPrefix = "Hello, ```

  We can now refactor our code

  ==CODE==

  After refactoring, re-reun your tests to make sure you haven't broken anything.

  Constants should improve performance of your application as it saves you creating the "Hello, " string
  instance every time Hello is called.

  To be clear, the performance boost is incredibly negligible for this example! But it's worth thinking about
  creating constants to capture the meaning of values and sometimes to aid performance.

# Hellow, world... again

  The next


