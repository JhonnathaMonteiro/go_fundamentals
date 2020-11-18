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

  ```const englishHelloPrefix = "Hello" ```

  We can now refactor our code

  ==CODE==

  After refactoring, re-reun your tests to make sure you haven't broken anything.

  Constants should improve performance of your application as it saves you creating the "Hello, " string
  instance every time Hello is called.

  To be clear, the performance boost is incredibly negligible for this example! But it's worth thinking about
  creating constants to capture the meaning of values and sometimes to aid performance.

# Hellow, world... again

  The next requirement is when our function is called with a empty string it defaults to printing "Hello, World",
  rather than "Hello, ".
  
  Start by writing a new failing test

  ==CODE==

  Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests
  around a "thing" and then have subsets describing different scenarios.

  A benefict of this approach is you can set up shared code that can be used in other tests.

  There is a repeated code when we check if the message is what we expect.

  Refactoring is not just for production code!

  It is important that your tests are clear specifications of what the code needs to do.

  We can and should refactor our tests.

  ==CODE==
  
  What have we done here ?

  We've refactored our  assertion into a function. This reduces duplication and improves readability of our tests. In Go
  You can declare functions inside another functions and assign them to variables. You can then call them, just like
  normal funtions. We need to pass in t *testing.T so that we can tell the test code to fail when we need to.

  t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line
  number reported will be in our function call rather than inside our test helper. This  will help other
  developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the
  output.

  Now that we hve a well-written failing test, let's fix the code, using an if  Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests
  around a "thing" and then have subsets describing different scenarios.

  A benefict of this approach is you can set up shared code that can be used in other tests.

  There is a repeated code when we check if the message is what we expect.

  Refactoring is not just for production code!

  It is important that your tests are clear specifications of what the code needs to do.

  We can and should refactor our tests.

  ==CODE==
  
  What have we done here ?

  We've refactored our  assertion into a function. This reduces duplication and improves readability of our tests. In Go
  You can declare functions inside another functions and assign them to variables. You can then call them, just like
  normal funtions. We need to pass in t *testing.T so that we can tell the test code to fail when we need to.

  t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line
  number reported will be in our function call rather than inside our test helper. This  will help other
  developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the
  output.

  Now that we hve a well-written failing test, let's fix the code, using an if.

  ==CODE==

  If we run our tests we should see it satisfies the new requeirement and we haven't accidentally broken the other 
  functionality.

# Back to source control

  Now we are happy with the code I would amend the previous commit so we only check in the lovely version of
  our code with its test.

## Discipline

  Let's go over the cycle again

  - Write a test
  - Make the complier pass
  - Run the test, see that it fails and check the error message is meaningful
  - Write enough code to make the test pass
  - Refactor

  On the face of it this may seens tedious but sticking to the feedback loop is important.

  Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with 
  the safety of tests.

  Seeing the test fails is an important check because it also lets you see what error message looks like.
  As a developer it can be very hard to work with a codebase when failing test do not give a clear idea as to
  what the problem is.

  By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of
  flow when writting your code.

  By not writing tests you are commiting to manually checking your code by running your sofware which
  breaks your state of flow and you won't be saving yourself any time, specially in the lon run.

  
# Keep going! More requirements

  Goodness me, we have more requirements. We now need to support a second paramenter, specifying the language
  of the greeting. If a languafe is passed in that we do not reconise, kust default to English.
  
  We should be confident that we can use TDD to flesh out this functionality easily!

  Write a test for a user passing in Spanish. Add it to the existing suite.

  ==CODE==
  
  Remember not to cheat! Test first. When you try and run the test, the compiler should complains because
  you are calling Hello with two arguments rather than one.

  ==CODE==

  Fix the compilation problems by adding another string arguments to Hello

  ==CODE==

  When you try and run the test again it will complain about not passing through enough arguments to
  Hello in your other tests and in hello.go

  ==CODE==

  Fix them by passing thtough empty strings. Now all your tests should compile and pass, apart from our
  scenario

  ==CODE==

  We can use if here to check the language is equal to "Spanish" and if so change the message

  ==CODE==

  The test should now pass.

  Now it is time to refactor. You should see some problems in the code, "magic" strings, some of which are
  repeated. Try and refactor it yourself with every change make sure you re-run the tests to make sure your 
  refactoring isn't breaking anything






















