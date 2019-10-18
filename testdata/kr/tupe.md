![tupe-cover](/Users/clay/go/src/github.com/tcd/md2pdf/testdata/kr/tupe-cover.png)

# Chapter 1: UNIX for Beginners

What is UNIX?
In the narrowest sense, it is a time-sharing operating system *kernel*: a program that controls the resources of a computer and allocates them among its users.
It lets users run their programs; it controls the peripheral devices (discs, terminals, printers, and the like) connected to the machine; and it provides a file system that manages there long-term storage of information such as programs, data, and documents.

In a broader sense, UNIX is often taken to include not only the kernel, but also essential programs like compilers, editors, command languages, programs for copying and printing files, and so on.

Still bore broadly, UNIX may even include programs developed by you or other users to be run on your system, such as tools for document preparation, routines for statistical analysis, and graphics packages.

Which of these uses of the name UNIX is correct depends on which level of the system you are considering.
When we use UNIX in the rest of this book, context should indicate which meaning is implied.

The UNIX system sometimes looks more difficult than it is - it's hard for a newcomer to know how to make the best use of the facilities available.
But fortunately it's not hard to get started - knowledge of only a few programs should get you off the ground.
This chapter is meant to help you to start using the system as quickly is possible.
It's an overview, not a manual; we'll cover most of the material again in more detail in later chapters.
We'll talk about these major areas:

- basics
  - logging in and out
  - simple commands
  - correcting typing mistakes
  - mail
  - inter-terminal communication
- day-to-day use
  - files and the file system
  - printing files
  - directories
  - commonly-used commands
- the command interpreter or *shell*
  - filename shorthands
  - redirecting input and output
  - pipes
  - setting erase and kill characters
  - defining your own search path for commands
  
If you've used a UNIX system before, most of this chapter should be familiar; you might want to skip straight to Chapter 2.

You will need a copy of the *UNIX Programmer's Manual*, even as you read this chapter; it's often easier for us to tell you to read about something in the manual than to repeat it's contents here.
This book is not supposed to replace it, but to show you how to make best use of the commands described in it.
Furthermore, there may be differences between what we say here and what is true on your system.
The manual has a permuted index at the beginning that's indispensable for finding the right programs to apply to a problem; learn to use it.

Finally, a word of advice: don't be afraid to experiment.
If you are a beginner, there are very few accidental things you can do to hurt yourself or other users.
So learn how things work by trying them.
This is a long chapter, and the best way to read it is a few pages at a time, trying things out as you go.

### 1.1 Getting started

#### Some prerequisites about terminals

To avoid explaining everything about using computers, we must assume you have some familiarity with computer terminals and how to use them.
If any of the following statements are mystifying, you should ask a local expert for help.

The UNIX system is *full duplex*: the characters you type on the keyboard are sent to the system, which sends them back to the terminal to be printed on the screen.
Normally, this *echo* process copies the characters directly to the screen, so you can see what you are typing, but sometimes, such as when you are typing a secret password, the echo is turned off so the characters do not appear on the screen.

Most of the keyboard characters are ordinary printing characters with no special significance, but a few tell the computer how to interpret your typing.
By far the most important of these is the RETURN key.
The RETURN key signifies the end of a line of input; the system echoes it by moving the terminal's cursor to the beginning of the line on the screen.
RETURN must be pressed before the system will interpret the characters you have typed.

RETURN is an example of a *control character* - an invisible character that controls some aspect of input and output on the terminal.
On any reasonable terminal, RETURN has a key of its own, but most control characters do not.
Instead, they must be typed by holding down the CONTROL key, sometimes called CTL or CNTL or CTRL, then pressing another key, usually a letter.
For example, RETURN may be typed by pressing down the RETURN key or, equivalently, holding down the CONTROL key and typing an 'm'.
RETURN might therefore be called a control-m, which we will write as `ctl-m`.
Other control characters include `ctl-d`, which tells a program that there is no more input; `ctl-g`, which rings the bell on the terminal; `ctl-h`, often called backspace, which can be used to correct typing mistakes; and `ctl-i`, often called tab, which advances the cursor to the next tab stop, much as on a regular typewriter.
Tab stops on UNIX systems are eight spaces apart.
Both the backspace and tab characters have their own keys on most terminals.

Two other keys have special meaning: DELETE, sometimes called RUBOUT or some abbreviation, and BREAK, sometimes called INTERRUPT.
On most UNIX systems, the DELETE key stops a program immediately, without waiting for it to finish.
On some systems, `ctl-c` provides this service.
And on some systems, depending on how the terminals are connected, BREAK is a synonym for DELETE or `ctl-c`.

#### A Session with UNIX

Let's begin with an annotated dialog between you and your UNIX system.
Throughout the examples in this book, what you type is printed in *`slanted letters`*, computer responses are in `typewriter-style characters`, and explanations are in *italics*.

```
Establish a connection: dial a phone or turn on a switch as necessary.
Your system should say
login: you
Password: 
You have mail.
$
$
$ date
Sun Sep 25 23:02:57 EDT 1983
$ who
jlb        tty0        Sep 25 13:59
you        tty2        Sep 25 23:01
mary       tty4        Sep 25 19:03
doug       tty5        Sep 25 19:22
egb        tty7        Sep 25 17:17
bob        tty8        Sep 25 20:48
$ mail
From doug Sun Sep 25 20:53 EDT 1983
give me a call sometime monday

?
From mary Sun Sep 25 19:07 EDT 1983
Lunch at noon tomorrow?

? d
$
$ mail mary
lunch at 12 is fine
ctl-d
$
```

Sometimes that's all there is to a session, though occasionally people do some work too.
The rest of this section will discuss the session above, plus other programs that make it possible to do useful things.

#### Logging in

You must have a login name and password, which you can get from your system administrator.
The UNIX system is capable of dealing with a wide variety of terminals, but it is strongly oriented towards devices with *lower case*; case distinctions matter!
If your terminal produces only upper case (like some video and portable terminals), life will be so difficult that you should look for another terminal.

Be sure the switches are set appropriately on your device, upper and lower case, full duplex, and any other settings that local experts advise, such as the speed, or *baud rate*.
Establish a connection using whatever magic is needed for your terminal; this may involve dialing a telephone or merely flipping a switch.
In either case, the system should type
```
login:
```
If it types garbage, you may be at the wrong speed; check the speed setting and other switches.
If that fails, press the BREAK or INTERRUPT key a few times, slowly.
If nothing produces a login message, you will have to get help.

When you get to the `login:` message, type your login name *in lower case*.
Follow it by pressing RETURN.
If a password is required, you will be asked for it, and printing will be turned off while you type it.

The culmination of your login efforts is a *prompt*, usually a single character, indicating that the system is ready to accept commands from you.
The prompt is most likely to be a dollar sign `$` or a percent sign `%`, but you can change it to anything you like; we'll show you how a little later.
The prompt is actually printed by a program called the *command interpreter* or *shell*, which is your main interface to the system.

There may be a message of the day just before the prompt, or a notification that you have mail.
You may also be asked what kind of terminal you are using; your answer helps the system to use any special properties the terminal might have.

#### Typing commands

Once you receive the prompt, you can type *commands*, which are requests that the system do something.
We will use *program* as a synonym for command.
When you see the prompt (let's assume it's `$`), type `date` and press RETURN.
The system should reply with the date and time, then print another prompt, so the whole transaction will look like this on your terminal:
```
$ date
Mon Sep 26 12:20:57 EDT 1983
$
```
Don't forget RETURN, and don't type the `$`.
If you think you're being ignored, press RETURN; something should happen.
RETURN won't be mentioned again, but you need it at the end of every line.

The next command to try is `who`, which tells you everyone who is currently logged in:
```
$ who
rlm        tty0        Sep 25 11:17
pjw        tty4        Sep 25 11:30
gerard     tty7        Sep 25 10:27
mark       tty9        Sep 25 07:59
you        ttya        Sep 25 12:20
```
The first column is the user name.
The second is the system's name for the connection being used ("tty" stands for "teletype", an archaic synonym for "terminal").
The rest tells when the user logged on.
You might also try:
```
$ whoami
you        ttya        Sep 25 12:20
```

If you make a mistake typing the name of a command, and refer to a nonexistent command, you will be told that no command of that name can be found:
```
$ whom
whom: not found
$
```
Of course, if you inadvertently type the name of an actual command, it will run, perhaps with mysterious results.

#### Strange terminal behavior

Sometimes your terminal will act strangely, for example, each letter may be typed twice, or RETURN may not put the cursor at the first column of the next line.
You can usually fix this by turning the terminal off and on, or by logging out and logging back in.
Or you can read the description of the command `stty` ("set terminal options") in Section 1 of the manual.
To get intelligent treatment of tab characters if your terminal doesn't have tabs, type the command
```
$ stty -tabs
```
and the system will convert tabs into the right number of spaces.
If your terminal does have computer-settable tab stops, the command `tabs` will set them correctly for you.
(You may actually have to say
```
$ tabs terminal-type
```
to make it work - see the `tabs` command description in the manual.)

#### Mistakes in typing

If you make a typing mistake, and see it before you have pressed RETURN, there are two ways to recover: *erase* characters one at a time or *kill* the whole line and re-type it.

If you type the *line kill* character, by default an at sign `@`, it causes the whole line to be discarded, just as if you'd never typed it, and starts you over on a new line:
```
$ ddtae@
date
Mon Sep 26 12:23:39 EDT 1983
```

The particular erase and line kill characters are *very* system dependent.
On many systems (including the one we use), the erase character has been changed to backspace, which works nicely on video terminals.
You can quickly check which is the case on your system:
```
$ datee⌫
datee⌫: not found
$ datee#
Mon Sep 26 12:26:08 EDT 1983
```
(We printed the backspace as ⌫ so you can see it.)
Another common choice is `ctl-u` for line kill.

We will use the sharp as the erase character for the rest of this section because it's visible, but make the mental adjustment if your system is different.
Later on, in "tailoring the environment," we will tell you how to set the erase and line kill characters to whatever you like, once and for all.

What if you must enter an erase or line kill character as part of the text?
If you precede either `#` or `@` by a backslash `\`, it loses its special meaning.
So to enter a `#` or `@`, type `\#` or `\@`.
The system may advance the terminal's cursor to the next line after your `@`, even if it was preceded by a backslash.
Don't worry - the at-sign has been recorded.

The backslash, sometimes called the *escape character*, is used extensively to indicate that the following character is in some way special.
To erase a backslash, you have to type two erase characters: `\##`.
Do you see why?

The characters you type are examined and interpreted by a sequence of programs before they reach their destination, and exactly how they are interpreted depends not only on where they end up but how they got there.

Every character you type is immediately echoed to the terminal, unless echoing is turned off, which is rare.
Until you press RETURN, the characters are held temporarily by the kernel, so typing mistakes can be corrected with the erase and line kill characters.
When an erase or line kill character is preceded by a backslash, the kernel discards the backslash and holds the following character without interpretation.

When you press RETURN, the characters being held are sent to the program that is reading from the terminal.
That program may in turn interpret the characters in special ways; for example, the shell turns off any special interpretation of a character if it is preceded by a backslash.
We'll come back to this in Chapter 3.
For now, you should remember that the kernel processes erase and line kill, and backslash only if it precedes erase or line kill; whatever characters are left after that may be interpreted by other programs as well.

#### Type-ahead

The kernel reads what you type as you type it, even if it's busy with something else, so you can type as fast as you want, whenever you want, even when some command is printing at you.
If you type while the system is printing, your input characters will appear intermixed with the other characters, but they will be stored away and interpreted in the correct order.
You can type commands one after another without waiting for them to finish or even to begin.

#### Stopping a program

You can stop most commands by typing the character DELETE.
The BREAK key found on most terminals may also work, although this is system dependent.
In a few programs, like text editors, DELETE stops whatever the program is doing but leaves you in that program.
Turning off the terminal or hanging up the phone will stop most programs.

If you just want output to pause, for example, to keep something critical from disappearing off the screen, type `ctl-s`.
The output will stop almost immediately; your program is suspended until you start it again.
When you want to resume, type `ctl-q`.

#### Logging out

The proper way to log out is to type `ctrl-d` instead of a command; this tells the shell that there is no more input.
(How this actually works will be explained in the next chapter.)
You can usually just turn off the terminal or hang up the phone, but whether this really logs you out depends on your system.

#### Mail

The system provides a postal system for communicating with other users, so some day when you log in, you will see this message
```
you have mail
```
before the first prompt.
To read your mail, type
```
$ mail
```
Your mail will be printed, one message at a time, most recent first.
After each item, `mail` waits for you to say what to do with it.
The two basic responses are `d`, which deletes the message, and RETURN, which does not (so it will still be there the next time you read your mail).
Other responses include `p` to reprint a message, `s` *filename* to save it in the file you named, and `q` to quit from `mail`.
(If you don't know what a file is, think of it as a place where you can store information under a name of your choice, and retrieve it later.
Files are the topic of Section 1.2 and indeed of much of this book.)

`mail` is one of those programs that is likely to differ from what we describe here; there are many variants.
Look in your manual for details.

Sending mail to someone is straightforward.
Suppose it is to go to the person with the login name `nico`. The easiest way is this:
```
$ mail nico
Now type in the text of the letter
on as many lines as you like ...
After the kast line of the letter
type a control-d
ctl-d
$
```
The `ctl-d` signals the end of the letter by telling the `mail` command that there is no more input.
If you change your mind half-way through composing the letter, press DELETE instead of `ctl-d`.
The half-formed letter will be stored in a file called `dead.letter` instead of being sent.

For practice, send mail to yourself, then type `mail` to read it.
(This isn't as aberrant as it might sound - it's a handy reminder mechanism.)

There are other ways to send mail - you can send a previously prepared letter, you can mail to a number of people all at once, and you may be able to send mail to people on other machines.
For more details see the description of the `mail` command in Section 1 of the *UNIX Programmer's Manual*.
Henceforth, we'll use the notation `mail`(1) to mean the page describing `mail` in Section 1 of the manual.
All of the commands discussed in this chapter are found in Section 1.

There may also be a calendar service (see `calendar`(1)); we'll show you in Chapter 4 how to set one up if it hasn't been done already.

#### Writing to other users

If your UNIX system has multiple users, someday, out of the blue, your terminal will print something like
```
Message from mary tty7...
```
accompanied by a startling beep.
Mary wants to write to you, but unless you take explicit action you won't be able to write back.
To respond, type 
```
$ write mary
```
This establishes a two-way communication path.
Now the lines that Mary types on her terminal will appear on yours and vice versa, although the path is slow, rather like talking to the moon.

If you are in the middle of something, you have to get to a state where you can type a command.
Normally, whatever program you are running has to stop or be stopped, but some programs, such as the editor and `write` itself, have a `!` command to escape temporarily to the shell - see Table 2 in Appendix 1.

The `write` command imposes no rules, so a protocol is needed to keep what you type from getting garbled up with what Mary types.
One convention is to take turns, ending each turn with `(o)`, which stands for "over," and to signal your intent to quit with `(oo)`, for "over and out."

#### News

Many UNIX systems provide a news service, to keep users abreast of interesting and not so interesting events.
Try typing
```
$ news
```
There is also a large network of UNIX systems that keep in tough through telephone calls; ask a local expert about `netnews` and USENET.

#### The manual

The *UNIX Programmer's Manual* describes most of what you need to know about the system.
Section 1 deals with commands, including those we discuss in this chapter.
Section 2 describes system calls, the subject of Chapter 7, and Section 6 has information about games.
The remaining sections talk about functions for use by C programmers, file formats, and system maintenance.
(The numbering of these sections varies from system to system.)
Don't forget the permuted index at the beginning; you can skim it quickly for commands that might be relevant to what you want to do.
There is also an introduction to the system that gives an overview of how things work.

Often the manual is kept on-line so that you can read it on your terminal.
If you get stuck on something, and can't find an expert to help, you can print any manual page on your terminal with the command `man command-name`.
Thus to read about the `who` command, type
```
$ man who
```
and, of course,
```
$ man man
```
tells about the `man` command.

#### Computer-aided instruction

Your system may have a command called `learn`, which provides computer-aided instruction on the file system and basic commands, the editor, document preparation, and even C programming.
Try
```
$ learn
```
If `learn` exists on your system, it will tell you what to do from there.
If that fails, you might also try `teach`.

#### Games

It's  not always admitted officially, but one of the best ways to get comfortable with a computer and a terminal is to play games.
The UNIX system comes with a modest supply of games, often supplemented locally.
Ask around, or see Section 6 of the manual.

### 1.2 Day-to-day use: files and common commands

Information in a UNIX system is stored in *files*, which are much like ordinary office files.
Each file has a name, contents, a place to keep it, and some administrative information such as who owns it and how big it is.
A file might contain a letter, or a list of names and addresses, or the source statements of a program, or data to be used by a program, or even programs in their executable form and other non-textual material.

The UNIX file system is organized so you can maintain your own personal files without interfering with files belonging to other people, and keep people from interfering with you too.
There are myriad programs that manipulate files, but for now, we will look at only the more frequently used ones.
Chapter 2 contains a systematic discussion of the file system, and introduces many of the other file-related commands.

#### Creating files - the editor

If you want to type a paper or letter or a program, how do you get the information stored in the machine?
Most of these tasks are done with a *text editor*, which is a program for storing and manipulating information in the computer.
Almost every UNIX system has a *screen editor*, an editor that takes advantage of modern terminals to display the effects of your editing changes in context as you make them.
Two of the most popular are `vi` and `emacs`.
We won't describe any specific screen editor here, however, partly because of typographic limitations, and partly because there is no standard one.

There is, however, an older editor called `ed` that is certain to be available on your system.
It takes no advantage of special terminal features, so it will work on any terminal.
It also forms the basis of other essential programs (including some screen editors), so it's worth learning eventually.
Appendix 1 contains a concise description.

No matter what editor you prefer, you'll have to learn it well enough to be able to create files.
We'll use `ed` here to make the discussion concrete, and to ensure that you can make our examples run on your system, but by all means use whatever editor you like best.

To use `ed` to create a file called `junk` with some text in it, do the following:
```
$ ed                            Invokes the text editor
a                               ed command to add text
now type in                     
whatever text you want ...
.                               Type a '.' by itself to stop adding text
w junk                          Write your text into a file called junk
39                              ed prints number of characters written
q                               Quit ed
$
```
The command `a` ("append") tells `ed` to start collecting text.
The "." that signals the end of the text must be typed at the beginning of a line by itself.
Don't forget it, for until it is typed, no other `ed` commands will be recognized - everything you type will be treated as text to be added.

The editor command `w` ("write") stores the information that you typed; `w junk` stores it in a file called `junk`.
The filename can be any word you like; we picked `junk` to suggest that this file isn't very important.

`ed` responds with the number of characters put in the file.
Until the `w` command, nothing is stored permanently, so if you hang up and go home the information is not stored in the file.
(If you hang up while editing, the data you were working on is saved in a file called `ed.hup`, which you can continue with at your next session.)
If the system crashes (i.e., stops unexpectedly because of software or hardware failure) while you are editing, your file will contain only what the last write command placed there.
But after `w` the information is recorded permanently; you can access it again later by typing
```
$ ed junk
```

Of course, you can edit the text you typed in, to correct spelling mistakes, change wording, rearrange paragraphs and the like.
When you're done, the `q` command ("quit") leaves the editor.

#### What files are out there?

Let's create two files, `junk` and `temp`, so we know what we have:
```
$ ed
a
To be or not to be
.
w junk
19
q
$ ed
a
That is the question.
.
w temp
22
q
$ 
```
The character counts from `ed` include the character at the end of each line, called *newline*, which is how the system represents RETURN.

The `ls` command lists the names (not contents) of files:
```
$ ls
junk temp
$
```
which are indeed the two files just created.
(There might be others as well that you didn't create yourself.)
The names are sorted into alphabetical order automatically.

`ls`, like most commands, has *options* that may be used to alter its default behavior.
Options follow the command name on the command line, and are usually made up of an initial minus sign (`-`) and a single letter meant to suggest the meaning.
For example, `ls -t` causes the files to be listed in "time" order: the order in which they were last changed, most recent first.
```
$ ls -t
temp
junk
$
```
The `-l` option gives a "long" listing that provides more information about each file:
```
$ls -l
total 2
-rw-r--r--  1 you          19 Sep 26 16:25 junk
-rw-r--r--  1 you          22 Sep 26 16:26 temp
```
`total 2` tells how many blocks of disk space the files occupy; a block is usually either 512 or 1024 characters.
The string `-rw-r--r--` tells who has permission to read and write the file; in this case, the owner (`you`) can read and write, but others can only read it.
The `1` that follows is the number of links to the file; ignore it until Chapter 2.
`you` is the owner of the file, that is, the person who created it.
`19` and `22` are the number of characters in the corresponding files, which agree with the numbers you got from `ed`.
The date and time tell when the file was last changed.

Options can be grouped: `ls -lt` gives the same data as `ls -l`, but sorted with most recent files first.
The `-u` option gives information on when files were used: `ls -lut` gives a long (`-l`) listing in the order of most recent use.
The option `-r` reverses the order of the output, so `ls -rt` lists in order of least recent use.
You can also name the files you're interested in, and `ls` will list the information about them only:
```
$ ls -l junk
-rw-r--r--  1 you          19 Sep 26 16:25 junk
$
```

The strings that follow the program name on the command line, such as `-l` and `junk` in the example above, are called the program's *arguments*.
Arguments are usually options or names of files to be used by the command.

Specifying options by a minus sign and a single letter, such as `-t` or the combined `-lt`, is a common convention.
In general, if a command accepts such optional arguments, they precede any filename arguments, but may otherwise appear in any order.
But UNIX programs are capricious in their treatment of multiple options.
For example, standard 7th Edition `ls` won't accept
```
$ ls -l -t
```
as a synonym for `ls -lt`, while other programs *require* multiple options to be separated.

As you learn more, you will find that there is little regularity or system to optional arguments.
Each command has its own idiosyncrasies, and its own choices of what letter means what (often different from the same function in other commands).
This unpredictable behavior is disconcerting and is often cited as a major flaw of the system.
Although the situation is improving - new versions often have more uniformity - all we can suggest is that you try to do better when you write your own programs, and in the meantime keep a copy of the manual handy.

#### Printing files - `cat`and `pr`

Now that you have some files, how do you look at their contents?
There are many programs to do that, probably more than are needed.
One possibility is to use the editor:
```
$ ed junk
19
1,$p
To be or not to be
q
$
```
`ed` begins by reporting the number of characters in `junk`; the command `1,$p` tells it to print all the lines in a file.
After you learn how to use the editor, you can be selective about the parts you print.

There are times when it's not feasible to use an editor for printing.
For example, there is a limit - several thousand lines - on how big a file `ed` can handle.
Furthermore, it will only print one file at a time, and sometimes you want to print several, one after another without pausing.
So here are a couple of alternatives.

First is `cat`, the simplest of all the printing commands.
`cat` prints the *contents* of all the files named by its arguments:
```
$ cat junk
To be or not to be
$ cat temp
That is the question.
$ cat junk temp
To be or not to be
That is the question.
```
The named file or files are catenated (hence the name `cat`) onto the terminal one after another with nothing between.
("Catenate" is a slightly obscure synonym for "concatenate")


#### Moving, copying, removing files - `mv`, `cp`, `rm`

Let's look at some other commands.
The first thing is to change the name of a file.
Renaming a file is done by "moving" it from one name to another, like this:
```
$ mv junk precious
```
This means that the file that used to be called `junk` is now called `precious`; the contents are unchanged.
When you run `ls` now, you will see a different list: `junk` is not there but `precious` is.
```
$ ls
precious
temp
$ cat junk
cat: can't open junk
$
```
Beware that if you move a file to another one that already exists, the target file is replaced.

To make a *copy* of a file (that is, to have two versions of something), use the `cp` command:
```
$ cp precious precious.save
```
makes a duplicate copy of `precious` in `precious.save`.

Finally, when you get tired of creating and moving files, the `rm` command removes all the files you name:
```
$ rm temp junk
rm: junk nonexistent
$
```
You will get a warning if one of the files to be removed wasn't there, but otherwise `rm`, like most UNIX commands, does its work silently.
There is no prompting or chatter, and error messages are curt and sometimes unhelpful.
Brevity can be disconcerting to newcomers, but experienced users find talkative commands annoying.

#### What's in a filename?

So far we have used filenames without ever saying what a legal filename is, so it's time for a couple of rules.
~First, filenames are limited to 14 characters.~
Second, although you can use almost any character in a filename, common sense says that you should stick to ones that are visible, and that you should avoid characters that might be used with other meanings.
We have already seen, for example, that in the `ls` command, `ls -t` means to list in time order.
So if you had a file whose name was `-t`, you would have a tough time listing it by name. 
(How would you do it?)
Besides the minus sign as a first character, there are other characters with special meaning.
To avoid pitfalls, you would do well to use only letters, numbers, the period, and the underscore until you're familiar with the situation.
(The period and the underscore are conventionally used to divide filenames into chunks, as in `precious.save` above.)
Finally, don't forget that case distinctions matter - `junk`, `Junk`, and `JUNK` are three different names.

#### A handful of useful commands

Now that you have the rudiments of creating files, listing their names, and printing their contents, we can look at a half-dozen file-processing commands.
To make the discussion concrete, we'll use a file called `poem` that contains a familiar verse by Augustus De Morgan.
Let's create it with `ed`:
```
$ ed
a
Great fleas have little fleas
	upon their backs to bite 'em.
And little fleas have lesser fleas,
	and so ad infinitum.
And the great fleas themselves, in turn,
	have greater fleas to go on;
While these again have greater still,
	and greater still and so on.
.
w poem
263
q
$
```

The first command counts lines, words, and characters in one or more files; it is named `wc` after its word-counting function:
```
$ wc poem
	8	46	258	poem.txt
$
```
That is, `poem` has 8 lines, 46 words, and 263 characters.
The definition of a "word" is very simple: any string of characters that doesn't contain a blank, tab, or newline.

`wc` will count more than one file for you (and print the totals), and it will also suppress any of the counts if requested.
See `wc`(1)

The second command is called `grep`; it searches files for lines that match a pattern.
(The name comes from the `ed` command `g/regular-expression/p`, which is explained in Appendix 1.)
Suppose you want to look for the word "fleas" in `poem`:
```
$ grep fleas poem
Great fleas have little fleas
And little fleas have lesser fleas,
And the great fleas themselves, in turn,
	have greater fleas to go on;
$
```
`grep` will also look for lines that *don't* match the pattern, when the option `-v` is used.
(It's named `v` after the editor command; you can think of it as inverting the sense of the match.)
```
$ grep -v fleas poem
	upon their backs to bite 'em.
	and so ad infinitum.
While these again have greater still,
	and greater still and so on.
$ 
```

`grep` can be used to search several files; in that case it will prefix the filename to each line that matches, so you can tell where the match took place.
There are also options for counting, numbering, and so on.
`grep` will also handle much more complicated patterns that just words like "fleas", but we will defer consideration of that until Chapter 4.

The third command is `sort`, which sorts its input into alphabetical order line by line.
This isn't very interesting for the poem, but let's do it anyway, just to see what it looks like:
```
$ sort poem
	and greater still and so on.
	and so ad infinitum.
	have greater fleas to go on;
	upon their backs to bite 'em.
And little fleas have lesser fleas,
And the great fleas themselves, in turn,
Great fleas have little fleas
While these again have greater still,
$
```
The sorting is line by line, but the default sorting order puts blanks first, then upper case letters, then lower case, so it's not strictly alphabetical.

`sort` has zillions of options to control the order of sorting - reverse order, numerical order, dictionary order, ignoring leading blanks, sorting on fields within the line, etc. - but usually one has to look up those options to be sure of them.
Here are a handful of the most common:

| option     | description                        |
|------------|------------------------------------|
| `sort -r`  | Reverse normal order               |
| `sort -n`  | Sort in numeric order              |
| `sort -nr` | Sort in reverse numeric order      |
| `sort -f`  | Fold upper and lower case together |
| `sort +n`  | Sort starting at n+1-st field      |

Chapter 4 has more information about `sort`.

Another file-examining command is `tail`, which prints the last 10 lines of a file.
That's overkill for our eight-line poem, but it's good for larger files.
Furthermore, tail has an option to specify the number of lines, so to print the last line of `poem`:
```
$ tail -1 poem
	and greater still, and so on.
$
```
`tail` can also be used to print a file starting at a specified line:
```
$ tail +3 filename
```
starts printing on the 3rd line.
(Notice the natural inversion of the minus sign convention for arguments.)

The final pair of commands is for comparing files.
Suppose that we have a variant of `poem` in the file `new_poem`:
```
$ cat poem
Great fleas have little fleas
	upon their backs to bite 'em.
And little fleas have lesser fleas,
	and so ad infinitum.
And the great fleas themselves, in turn,
	have greater fleas to go on;
While these again have greater still,
	and greater still and so on.
$ cat new_poem
Great fleas have little fleas
	upon their backs to bite them.
And little fleas have lesser fleas,
	and so on ad infinitum.
And the great fleas themselves, in turn,
	have greater fleas to go on;
While these again have greater still,
	and greater still and so on.
```
There's not much difference between the two files; in fact you'll have to look hard to find it.
This is where comparison commands come in handy.
`cmp` finds the first place where two files differ:
```
$ cmp poem new_poem
poem new_poem differ: char 58, line 2
$
```
This says that the files are different in the second line, which is true enough, but it doesn't say what the difference is, not does it identify any differences beyond the first.

The other file comparison command is `diff`, which reports on all lines that are changed, added, or deleted:
```
$ diff poem new_poem
2c2
< 	upon their backs to bite 'em.
---
> 	upon their backs to bite them.
4c4
< 	and so ad infinitum.
---
> 	and so on ad infinitum.
$
```
This says that line 2 in the first file (`poem`) has to be changed into line 2 of the second file (`new_poem`), and similarly for line 4.

Generally speaking, `cmp` is used when you want to be sure that two files really have the same contents.
It's fast and works on any kind of file, not just text.
`diff` is used when files are expected to be somewhat different, and you want to know exactly what lines differ.
`diff` works only on files of text.

#### A summary of file system commands

Table 1.1 is a brief summary of the commands we've seen so fat that deal with files.

##### Table 1.1

| command                       | summary                                                     |
|-------------------------------|-------------------------------------------------------------|
| `ls`                          | list names of all files in current directory                |
| `ls` *filenames*              | list only the named files                                   |
| `ls -t`                       | list in time order, most recent first                       |
| `ls -l`                       | list long: more information; also `ls -lt`                  |
| `ls -u`                       | list by last time used; also `ls -lu`, `ls -lut`            |
| `ls -r`                       | list in reverse order; also `-rt`, `rlt`, etc.              |
| `ed` *filename*               | edit named file                                             |
| `cp` *file1 file2*            | copy *file1* to *file2*, overwrite old *file2* if it exists |
| `mv` *file1 file2*            | move *file1* to *file2*, overwrite old *file2* if it exists |
| `rm` *filenames*              | remove named files, irrevocably                             |
| `cat` *filenames*             | print contents of named files                               |
| `wc` *filenames*              | count lines, words, and characters for each file            |
| `wc -l` *filenames*           | count lines for each file                                   |
| `grep` *pattern filenames*    | print lines matching *pattern*                              |
| `grep -v` *pattern filenames* | print lines not matching *pattern*                          |
| `sort` *filenames*            | sort files alphabetically by line                           |
| `tail` *filename*             | print last 10 lines of file                                 |
| `tail -n` *filename*          | print last *n* lines of file                                |
| `tail +n` *filename*          | start printing file at line *n*                             |
| `cmp` *file1 file2*           | print location of first difference                          |
| `diff` *file1 file2*          | print all differences between files                         |

### 1.3 More about files: directories

The system distinguishes your file called `junk` from anyone else's of the same name.
The distinction is made by grouping files into *directories*, rather in the way that books are placed on shelves in a library, so files in different directories can have the same name without any conflict.

Generally each user has a personal or *home directory*, sometimes called login directory, that contains only the files that belong to him or her.
When you log in, you are "in" your home directory.
You may change the directory you are working in - often called your working or *current directory* - but your home directory is always the same.
Unless you take special action, when you create a new file it is made in your current directory.
Since this is initially your home directory, the file is unrelated to a file of the same name that might exist in someone else's directory.

A directory can contain other directories as well as ordinary files ("Great directories have lesser directories...").
The natural way to picture this organization is as a tree of directories and files.
It is possible to move around within this tree, and to find any file in the system by starting at the root of the tree and moving along the proper branches.
Conversely, you can start where you are and move toward the root.

Let's try the latter first.
Our basic tool is the command `pwd` ("print working directory"), which prints the name of the directory you are currently in:
```
$ pwd
/usr/you
$
```
This says that you are currently in the directory `you`, in the directory `usr`, which in turn is in the *root directory*, which is conventionally just `/`.
The `/` characters separate the components of the name; the limit of 14 characters mentioned above applies to each component of such a name.
On many systems, `/usr` is a directory that contains the directories of all the normal users of the system.
(Even if your home directory is not `/usr/you`, `pwd` will print something analogous, so you should be able to follow what happens below.)

Now if you type
```
$ ls /usr/you
```
you should get exactly the same list of file names as you get from a plain `ls`.
When no arguments are provided, `ls` lists the contents of the current directory; given the name of a directory, it lists the contents of that directory.

Next, try
```
$ ls /usr
```
This should print a long series of names, among which is your own login directory `you`.

The next step is to try listing the root itself.
You should get a response similar to this:
```
$ ls /
bin
boot
dev
etc
lib
tmp
unix
usr
$
```
(Don't be confused by the two meanings of `/`: it's both the name of the root directory and a separator in filenames.)
Most of these are directories, but `unix` is actually a file containing the executable form of the UNIX kernel.
More of this in Chapter 2.

Now try
```
$ cat /usr/you/junk
```
(if `junk` is still in your directory).
The name
```
/usr/you/junk
```
is called the *pathname* of the file.
"Pathname" has an intuitive meaning: it represents the full name of the path from the root through the tree of directories to a particular file.
It is a universal rule in the UNIX system that wherever you can use an ordinary filename, you can use a pathname.

The file system is structured like a genealogical tree; here is a picture that may make it clearer.
```mermaid
graph TD
A[/]
A ---B1[bin]
A ---B2[boot]
A ---B3[dev]
A ---B4[etc]
A ---B5[lib]
A ---B6[tmp]
A ---B7[unix]
A ---B8[usr]
B8---C1[you]
B8---C2[mike]
B8---C3[paul]
B8---C4[mary]
C1---D1[junk]
C3---D2[junk]
C3---D3[temp]
C4---D4[junk]
C4---D5[temp]
```
Your file named `junk` is unrelated to Paul's or to Mary's.

Pathnames aren't too exciting if all the files of interest are in your own directory, but if you work with someone else or on several projects concurrently, they become handy indeed.
For example, your friends can print your `junk` by saying
```
$ cat /usr/you/junk
```
Similarly, you can find out what files Mary has by saying
```
$ ls /usr/mary
data
junk
$
```
or make your own copy of one of her files by
```
$ cp /usr/mary/data data
```
or edit her file:
```
$ ed /usr/mary/data
```

If Mary doesn't want you poking around in her files, or vice versa, privacy can be arranged.
Each file and directory has read-write-execute permissions for the owner, a group, and everyone else, which can be used to control access.
(Recall `ls -l`.)
In our local systems, most users most of the time find openness of more benefit than privacy, but policy may be different on your system, so we'll get back to this in Chapter 2.

As a final set of experiments with pathnames, try
```
$ ls /bin /usr/bin
```
Do some of the names look familiar?
When you run a command by typing its name after the prompt, the system looks for a file of that name.
It normally looks first in your current directory (where it probably doesn't find it), then in `/bin`, and finally in `/usr/bin`.
There is nothing special about commands like `cat` or `ls`, except that they have been collected into a couple of directories to be easy to find and administer.
To verify this, try to execute some of these programs by using their full pathnames:
```
$ /bin/date
Mon Sep 26 23:29:32 EDT 1983
$ /bin/who
srm	tty1	Sep 26 22:20
cvw	tty4	Sep 26 22:40
you	tty5	Sep 26 23:04
$
```

Try
```
$ ls /usr/games
```
and do whatever comes naturally.
Things might be more fun outside of normal working hours.

#### Changing directory - `cd`

If you work regularly with Mary on information in her directory, you can say "I want to work on Mary's files instead of my own."
This is done by changing your current directory with the `cd` command:
```
$ cd /usr/mary
```
Now when you use a filename (without /'s) as an argument to `cat` or `pr`, it refers to the file in Mary's directory.
Changing directories doesn't affect any permissions associated with a file - if you couldn't access a file from your own directory, changing to another directory won't alter that fact.

It is usually convenient to arrange your own files so that all files related to one thing are in a directory separate from other projects.
For example, if you want to write a book, you might want to keep all the text in a directory called `book`.
The command `mkdir` makes a new directory.
```
$ mkdir book         Make a directory
$ cd book            Go to it
$ pwd                Make sure you're in the right place
/usr/you/book        
...                  Write the book (several minutes pass)
$ cd ..              Move up one level in the file system
$ pwd
/usr/you
$
```
`..` refers to the parent of whatever directory you are currently in, the directory one level closer to the root.
`.` is a synonym for the current directory.
```
$ cd        Return to the home directory
```
all by itself will take you back to your home directory, the directory where you log in.

Once your book is published, you can clean up the files.
To remove the directory `book`, remove all the files in it (we'll show you a fast way shortly), then `cd` to the parent directory of `book` and type
```
$ rmdir book
```
`rmdir` will only remove an empty directory

### 1.4 The Shell

When the system prints the prompt `$` and you type commands that get executed, it's not the kernel that is talking to you, but a go-between called the command interpreter or *shell*.
The shell is just an ordinary program like `date` or `who`, although it can do some remarkable things.
The fact that the shell sits between you and the facilities of the kernel has real benefits, some of which we'll talk about here.
There are three main ones:
- Filename Shorthands
  - You can pick up a whole set of filenames as arguments to a program by specifying a pattern for the names - the shell will find the filenames that match your pattern. 
- Input-output redirection
  - You can arrange for the output of any program to go into a file instead of the terminal, and for the input to come from a file instead of the terminal.  Input and output can even be connected to other programs.
- Personalizing the environment
  - You can define your own commands and shorthands

#### Filename shorthand

Let's begin with filename patterns.
Suppose you're typing a large document like a book.
Logically this divides into many small pieces, like chapters and perhaps sections.
Physically it should be divided too, because it is cumbersome to edit large files.
Thus you should type the document as a number of files.
You might have separate files for each chapter, called `ch1`, `ch2`, etc.
Or, if each chapter were broken into sections, you might create files called
```
ch1.1
ch1.2
ch1.3
...
ch2.1
ch2.2
...
```
which is the organization we used for this book.
With a systematic naming convention, you can tell at a glance where a particular file fits into the whole.

What if you want to print the whole book?
You could say
```
$ pr ch1.1 ch1.2 ch1.3 ...
```
but you would soon get bored typing filenames and start to make mistakes.
This is where filename shorthand comes in.
If you say
```
$ pr ch*
```
the shell takes the `*` to mean "any string of characters," so `ch*` is a pattern that matches all filenames in the current directory that begin with `ch`.
The shell creates the list, in alphabetical order, and passes the list to `pr`.
The `pr` command never sees the `*`; the pattern match that the shell does in the current directory generates a list of strings that are passed to `pr`.

The crucial point is that filename shorthand is not a property of the `pr` command, but a service of the shell.
Thus you can use it to generate a sequence of filenames for *any* command.
For example, to count the words in the first chapter:
```
$ wc ch1.*
 113	  562	3200 	ch1.0
 935	 4081	22435	ch1.1
 974	 4191	22756	ch1.2
 378	 1561	8481 	ch1.3
1293	 5298	28841	ch1.4
  33	  194	1190 	ch1.5
  75	  323	2030 	ch1.6
3801	16210	88933	total
$
```

There is a program called `echo` that is especially valuable for experimenting with the meaning of the shorthand characters.
As you might guess, `echo` does nothing more than echo its arguments:
```
$ echo hello world
hello world
$
```
But the arguments can be generated by pattern-matching:
```
$ echo ch1.*
```
lists all of the names of all the files in Chapter 1,
```
$ echo *
```
lists *all* the filenames in the current directory in alphabetical order,
```
$ pr *
```
prints all your files (in alphabetical order), and
```
$ rm *
```
removes *all files* in your current directory.
(You had better be *very* sure that's what you wanted to say!)

The `*` is not limited to the last position in a filename - `*`s can be anywhere and can occur several times.
Thus
```
$ rm *.save
```
removes all files that end with `.save`.

Notice that the filenames are sorted alphabetically, which is not the same as numerically.
If your book has ten chapters, the order might not be what you intended, since `ch10` comes before `ch2`:
```
$ echo *
ch1.1 ch1.2 ... ch10.1 ch10.2 ... ch2.1 ch2.2
$
```

The `*` is not the only pattern-matching feature provided by the shell, although it's by far the most frequently used.
The pattern `[...]` matches any of the characters inside the brackets.
A range of consecutive letters or digits can be abbreviated:
```
$ pr ch[12346789]*        Print chapters 1,2,3,4,6,7,8,9, but not 5
$ pr ch[1-46-9]*          Same thing
$ rm temp[a-z]            Remove any of tempa, ..., tempz that exist
```
The `?` pattern matches any single character:
```
$ ls ?                    List files with single
$ ls -l ch?.1             List ch1.1, ch2.1, ch3.1, etc. but not ch10.1
$ rm temp?                Removes files temp1, ..., tempa, etc.
```
Note that the patterns match only *existing* filenames.
In particular, you cannot make up new filenames by using patterns.
For example, if you wanted to expand `ch` to `chapter` in each filename, you cannot do it this way:
```
$ mv ch.* chapter.*       Doesn't work!
```
because `chapter.*` matches no existing filenames.

Pattern characters like `*` can be used in pathnames as well as simple filenames; the match is done for each component of the path that contains a special character.
Thus `/usr/mary/*` performs the match in `/usr/mary`, and `/usr/*/calendar` generates a list of pathnames of all user `calendar` files.

If you should ever have to turn off the special meaning of `*`, `?`, etc., enclose the entire argument in single quotes, as in
```
$ ls '?'
```
You can also precede a special character with a backslash:
```
$ ls \?
```
(Remember that because `?` is not the erase or line kill character, this backslash is interpreted by the shell, not the kernel.)
Quoting is treated at length in Chapter 3.

#### Input-output redirection

Most of the commands we have seen so far produce output to the terminal; some, like the editor, also take their input from the terminal.
It is nearly universal that the terminal can be replaced by a file for either or both of input or output.
As one example:
```
$ ls
```
makes a list of filenames on your terminal.
But if you say
```
$ ls > filelist
```
that same list of filenames will be placed in the file `filelist` instead.
The symbol `>` means "put the output in the following file, rather than on the terminal."
The file will be created if it doesn't already exist, or the previous contents overwritten if it does.
Nothing is produced on your terminal.
As another example, you can combine several files into one by capturing the output of `cat` in a file.
```
$ cat f1 f2 f3 > temp
```
The symbol `>>` operates much as `>` does, except that it means "add to the end of."
That is,
```
$ cat f1 f2 f3 >> temp
```
copies the contents of `f1`, `f2`, and `f3` onto the end of whatever is already in `temp`, instead of overwriting the existing contents.
As with `>`, if `temp` doesn't exist, it will be created initially empty for you.

In a similar way, the `<` symbol means to take the input for a program from the following file, instead of from the terminal.
Thus, you can prepare a letter in file `let`, then send it to several people with
```
$ mail mary joe tom bob < let
```
In all of these examples, blanks are optional on either side of `>` or `<`, but our formatting is traditional.

Given the capability of redirecting output with `>`, it becomes possible to combine commands to achieve effects not possible otherwise.
For example, to print an alphabetical list of users,
```
$ who > temp
$ sort < temp
```
Since `who` prints one line of output per logged-on user, and `wc -l` counts lines (suppressing the word and character counts), you can count users with
```
$ who > temp
$ wc -l < temp
```
You can count the files in the current directory with
```
$ ls > temp
$ wc -l < temp
```
though this includes the filename `temp` itself in the count.
You can print the filenames in three columns with
```
$ ls > temp
$ pr -3 < temp
```
And you can see if a particular user is logged on by combining `who` and `grep`:
```
$ who > temp
$ grep mary < temp
```

In all of these examples, as with filename pattern characters like `*`, it's important to remember that interpretation of `<` and `>` is being done by the shell, not by the individual programs.
Centralizing the facility in the shell means that input and output redirection can be used with any program; the program itself isn't aware that something unusual has happened.

This brings up an important convention.
The command
```
$ sort < temp
```
sorts the contents of the file `temp`, as does
```
$ sort temp
```
but there is a difference.
Because the string `< temp` is interpreted by the shell, `sort` does not see the filename `temp` as an argument; instead it sorts its *standard input*, which the shell has redirected so it comes from the file.
The latter example, however, passes the name `temp` as an argument to `sort`, which reads the file and sorts it.
`sort` can be given a list of filenames, as in
```
$ sort temp1 temp2 temp3
```
but if no filenames are given, it sorts its standard input.
This is an essential property of most commands: if no filenames are specified, the standard input is processed.
This means that you can simply type at commands to see how they work.
For example,
```
$ sort
ghi
abc
def
ctl-d
abc
def
ghi
$
```
In the next section, we will see how this principle is exploited.

#### Pipes

All of the examples at the end of the previous section rely on the same trick: putting the output of one program into the input if another via a temporary file.
But the temporary file has no other purpose; indeed, it's clumsy to have to use such a file.
This observation leads to one of the fundamental contributions of the UNIX system, the idea of a *pipe*.
A pipe is a way to connect the output of one program to the input of another program without any temporary file; a *pipeline* is a connection of two or more programs through pipes.

Let us revise some of the earlier examples to use pipes instead of temporary files.
The vertical bar character `|` tells the shell to set up a pipeline:
```
$ who | sort             Print sorted list of users
$ who |  wc -l           Count users
$ ls | wc -l             Count files
$ ls | pr -3             3-column list of filenames
$ who | grep mary        Look for particular user
```

Any program that reads from the terminal can read from a pipe instead; any program that writes on the terminal can write to a pipe.
This is where the convention of reading the standard input when no files are named pays off: any program that adheres to the convention can be used in pipelines.
`grep`, `pr`, `sort`, and `wc` are all used that way in the pipelines above.

You can have as many programs in a pipeline as you wish:
```
$ ls | pr -3 | lpr
```
creates a 3-column list of filenames on the line printer, and
```
$ who | grep mary | wc -l
```
counts how many times Mary is logged in.

The programs in a pipeline actually run at the same time, not one after another.
This means that the programs in a pipeline can be interactive; the kernel looks after whatever scheduling and synchronization is needed to make it all work.

As you probably suspect by now, the shell arranges things when you ask for a pipe; the individual programs are oblivious to the redirection.
Of course, programs have to operate sensibly if they are to be combined this way.
Most commands follow a common design, so they will fit properly into pipelines at any position.
Normally a command invocation looks like
```
command        optional-arguments        optional-filenames
```
If no filenames are given, the commands reads its standard input, which is by default the terminal (handy for experimenting) but which can be redirected to come from a file or a pipe.
At the same time, on the output side, most commands write their output to the *standard output*, which is by default sent to the terminal.
But it too can be redirected to a file or a pipe.

Error messages from commands have to be handled differently, however, or they might disappear into a file or down a pipe.
So each command has a *standard error* output as well, which is normally directed to your terminal.
Or, as a picture:
```mermaid
graph LR
A[Standard Input or files] -->B{command, options}
B-->C[Standard Output]
B-->D[Standard Error]
```
Almost all of the commands we have talked about so far fit this model; the only exceptions are commands like `date` and `who` that read no input, and a few like `cmp` and `diff` that have a fixed number of file inputs.
(But look at the `-` option on these.)

#### Processes

The shell does quite a few things besides setting up pipes.
Let us turn briefly to the basics of running more than one program at a time, since we have already seen a bit of that with pipes.
For example, you can run two programs with one command line by separating the commands with a semicolon; the shell recognizes the semicolon and breaks the line into two commands:
```
$ date; who
Tue Sep 27 01:03:17 EDT 1983
ken
dmr
rob
bwk
kk
you
ber
$
```
Both commands are executed (in sequence) before the shell returns with a prompt character.

You can also have more than one program running simultaneously if you wish.
For example, suppose you want to do something time-consuming like counting the words in your book, but you don't want to wait for `wc` to finish before you start something else.
Then you can say
```
$ wc ch* > wc.out &        
6994                       Process-id printed by the shell
$
```
The ampersand `&` at the end of a command line says to the shell "start this command running, then take further commands from the terminal immediately," that is, don't wait for it to complete.
Thus the command will begin, but you can do something else while it's running.
Directing the output into the file `wc.out` keeps it from interfering with whatever you're doing at the same time.

An instance of a running program is called a *process*.
The number printed by the shell for a command initiated with `&` is called the *process-id*; you can use it in other commands to refer to a specific running program.

It's important to distinguish between programs and processes.
`wc` is a program; each time you run the program `wc`, that creates a new process.
If several instances of the same program are running at the same time, each is a separate process with a different process-id.

If a pipeline is initiated with `&`, as in
```
$ pr ch* | lpr &
6951                       Process-id of lpr 
$
```
the processes in it are all started at once - the `&` applies to the whole pipeline.
Only one process-id is printed, however, for the last process in the sequence.

The command
```
$ wait
```
waits until all processes initiated with `&` have finished.
If it doesn't return immediately, you have commands still running.
You can interrupt `wait` with DELETE.

You can use the process-id printed by the shell to stop a process initiated with `&`:
```
$ kill 6944
```
If you forget the process-id, you can use the command `ps` to tell you about everything you have running.
If you are desperate, `kill 0` will kill all your processes except your login shell.
And if you're curious about what other users are doing, `ps -ag` will tell you about *all* processes that are currently running.
Here is some sample output:
```
$ ps -ag
 PID TTY TIME CMD
  36 co  6:29 /etc/cron
6423 5   0:02 -sh
6704 1   0:04 -sh
6722 1   0:12 vi paper
4430 2   0:03 -sh
6612 7   0:03 -sh
6628 7   1:13 rogue
6843 2   0:02 write dmr
6949 4   0:01 login bimmler
6952 5   0:08 pr ch1.1 ch1.2 ch1.3 ch1.4
6951 5   0:03 lpr
6959 5   0:02 ps -ag
6844 1   0:02 write rob
$
```
PID is the process id; TTY is the terminal associated with the process (as in `who`); TIME is the processor time used in minutes and seconds; and the rest is the command being run.
`ps` is one of those commands that is different on different versions of the system, so your output may not be formatted like this.
Even the arguments may be different - see the manual page for `ps`(1).

Processes have the same sort of hierarchical structure that files do: each process has a parent, and may well have children.
Your shell was created by a process associated with whatever terminal line connects you to the system.
As you run commands, those processes are the direct children of your shell.
If you run a program from within one of those, for example with the `!` command to escape from `ed`, that creates its own child process which is thus a grandchild of the shell.

Sometimes a process takes so long that you would like to start it running, then turn off the terminal and go home without waiting for it to finish.
But if you turn off your terminal or break your connection, the process will normally be killed even if you used `&`.
The command `nohup` ("no hangup") was created to deal with this situation: if you say
```
$ nohup command &
```
the command will continue to run if you log out.
Any output from the command is saved in a file called `nohup.out`.
There is no way to `nohup` a command retroactively.

If your process will take a lot of processor resources, it is kind to those who share your system to run your job with lower than normal priority; this is done by another program called `nice`:
```
$ nice expensive-command &
```
`nohup` automatically calls `nice`, because if you're going to log out you can afford to have the command take a little longer.

Finally, you can simply tell the system to start your process at some wee hour of the morning when normal people are asleep, not computing.
The command is called `at`(1):
```
$ at time
whatever commands
you want
ctl-d
$
```
This is the typical usage, but of course the commands could come from a file:
```
$ at 3am < file
$
```
Times can be written in 24-hour style like `2130`, or 12 hour style like `930pm`.

#### Tailoring the environment

One of the virtues of the UNIX system is that there are several ways to bring it closer to your personal taste or the conventions of your local computing environment.
For example, we mentioned earlier the problem of different standards for the erase and line kill characters, which by default are usually `#` and `@`.
You can change these any time you want with
```
$ stty erase e kill k
```
where `e` is whatever character you want for erase and `k` is for line kill.
But it's a bother to have to type this every time you log in.

The shell comes to the rescue.
If there is a file named `.profile` in your login directory, the shell will execute the commands in it when you log in, before printing the first prompt.
So you can put commands into `.profile` to set up your environment as you like it, and they will be executed every time you log in.

The first thing most people put in their `.profile` is
```
stty erase ⌫
```
We're using `⌫` here so you can see it, but you could put a literal backspace in your `.profile`.
`stty` also understands the notation `^x` for `ctl-x`, so you can get the same effect with
```
stty erase '^h'
```
because `ctl-h` is backspace.
(The `^` character is an obsolete synonym for the pipe operator `|`, so you must protect it with quotes.)

If your terminal doesn't have sensible tab stops, you can add `-tabs` to the `stty` line
```
stty erase '^h' -tabs
```

If you like to see how busy the system is when you log in, add
```
who | wc -l
```
to count the users.
If there's a news service, you can add `news`.
Some people like a fortune cookie:
```
/usr/games/fortune
```
after a while you may decide that it is taking too long to log in, and cut your `.profile` back to the bare necessities.

Some of the properties of the shell are actually controlled by so-called *shell-variables*, with values that you can access and set yourself.
For example, the prompt string, which we have been showing as `$`, is actually stored in a shell variable called `PS1`, and you can set it to anything you like, like this:
```
PS1='Yes dear? '
```
The quotes are necessary since there are spaces in the prompt string.
Spaces are not permitted around the `=` in this construction.

The shell also treats the variables `HOME` and `MAIL` specially.
`HOME` is the name of your home directory; it is normally set properly without having to be in `.profile`.
The variable `MAIL` names the standard file where your mail is kept.
If you define it for the shell, you will be notified after each command if new mail has arrived:
```
MAIL=/usr/spool/mail/you
```
(The mail file may be different on your system; `/usr/mail/you` is also common.)

Probably the most useful shell variable is the one that controls where the shell looks for commands.
Recall that when you type the name of a command, the shell normally looks for it first in the current directory, then in `/bin`, and then in `/usr/bin`.
This sequence of directories is called the *search path*, and is stored in a shell variable called `PATH`.
If the default search path isn't what you want, you can change it, again usually in your `.profile`.
For example, this line sets the path to the standard one plus `/usr/games`:
```
PATH=.:/bin:/usr/bin:/usr/games
```
The syntax is a bit strange: a sequence of directory names separated by colons.
Remember that `.` is the current directory.
You can omit the `.`; a null component in `PATH` means the current directory.

An alternate way to set `PATH` in this specific case is to simply augment the previous value:
```
PATH=$PATH:/usr/games
```
You can obtain the value of any shell variable by prefixing its name with a `$`.
In the example above, the expression `$PATH` retrieves the current value, to which the new part is added, and the result is assigned back to `PATH`.
You can verify this with `echo`:
```
$ echo PATH is $PATH
PATH is :/bin:/usr/bin:/usr/games
$ echo HOME is $HOME
HOME is /usr/you
$
```
If you have some of your own commands, you might want to collect them in a directory of your own and add that to your search path as well.
If you have some of your own commands, you might want to collect them in a directory of your own and add that to your search path as well.
In that case, your `PATH` might look like this:
```
PATH=:$HOME/bin:/bin:/usr/bin:/usr/games
```
We'll talk about writing your own commands in Chapter 3.

Another variable, often used by text editors fancier than `ed`, is `TERM`, which names the kind of terminal you are using.
That information may make it possible for programs to manage your screen more effectively.
Thus you might add something like
```
TERM=adm3
```
ti your `.profile` file.

It is also possible to use variables for abbreviation.
If you find yourself frequently referring to some directory with a long name, it might be worthwhile adding a line like
```
d=/horribly/long/directory/name
```
to your profile, so that you can say things like
```
$ cd $d
```
Personal variables like `d` are conventionally spelled in lower case to distinguish them from those used by the shell itself, like `PATH`.

Finally, it's necessary to tell the shell that you intend to use the variables in other programs; this is done with the command `export`, to which we will return in Chapter 3:
```
export MAIL PATH TERM
```

To summarize, here is what a typical `.profile` file might look like:
```
$ cat ~/.profile
stty erase '^h' -tabs
export MAIL=/usr/spool/mail/you
export PATH=:$HOME/bin:/bin:/usr/bin:/usr/games
export TERM=adm3
export b=$HOME/book
date
who | wc -l
```

We have by no means exhausted the services the shell provides.
One of the most useful is that you can create your own commands by packaging existing commands into a file to be processed by the shell.
It is remarkable how much can be achieved by this fundamentally mechanism.
Our discussion of it begins in Chapter 3.

### 1.5 The rest of the UNIX system

There's much more to the UNIX system than we've addressed in this chapter, but then there's much more to this book.
By now, you should feel comfortable with the system and, particularly, with the manual.
When you have specific questions about when or how to use commands, the manual is the place to look.

It is also worth browsing the manual occasionally, to refresh your knowledge of familiar commands and to discover new ones.
The manual describes many programs we won't illustrate, including compilers for languages like FORTRAN 77; calculator programs such as `bc`(1); `cu`(1) and `uucp`(1) for inter-machine communication; graphics packages; statistics programs; and esoterica such as `units`(1).

As we've said before, this book does not replace the manual, it supplements it.
In the chapters that follow, we will look at pieces and programs of the UNIX system, starting from the information in the manual but following the threads that connect the components.
Although the program interrelationships are never made explicit in the manual, they form the fabric of the UNIX programming environment.


# Chapter 2: The File System

Everything in the UNIX system is a file.
That is less of an oversimplification than you might think.
When the first version of the system was being designed, before it even had a name, the discussions focused on the structure of a file system that would be clean and easy to use.
The file system is central to the success and convenience of the UNIX system.
It is one of the best examples of the "keep it simple" philosophy, showing the power achieved by careful implementation of a few well-chosen ideas.

To talk comfortably about commands and their interrelationships, we need a good background in the structure and outer workings of the file system.
This chapter covers most of the details of using the file system:
- what files are
- how they are represented
- directories and the file system hierarchy permissions
- inodes (the system's internal record of files)
- device files
Because most use of the UNIX system deals with manipulating files, there are many commands for file investigation or rearrangement; this chapter introduces the more commonly used ones.

### 2.1 The basics of files

A file is a sequence of bytes.
(A byte is a small chunk of information, typically 8 bits long. For our purposes, a byte is equivalent to a character.)
No structure is imposed on a file by the system, and no meaning is attached to its contents - the meaning of the bytes depends solely on the programs that interpret the file.
Furthermore, as we shall see, this is true not just of disc files but of peripheral devices as well.
Magnetic tapes, mail messages, characters typed on the keyboard, line printer output, data flowing in pipes - each of these files is just a sequence of bytes as far as the system and the programs in it are concerned.

The best way to learn about files is to play with them, so start by creating a small file:
```
$ ed
a
now is the time
for all good people
.
w junk
36
q
$ ls -l junk
-rw-r--r--  1  you     36 Sep 27 06:11 junk
$
```
`junk` is a file with 36 bytes - the 36 characters you typed while appending (except, of course, for correction of any typing mistakes).
To see the file,
```
$ cat junk
now is the time
for all good people
$ 
```
`cat` shows what the file looks like.
The command `od` (octal dump) prints a visible representation of all the bytes of a file:
```
$ od -c junk
0000000   n   o   w       i   s       t   h   e       t   i   m   e  \n
0000020   f   o   r       a   l   l       g   o   o   d       p   e   o
0000040   p   l   e  \n
0000044
$
```
The `-c` option means "interpret bytes as characters."
Turning on the `-b` option will show the bytes as octal (base 8) numbers as well:
```
$ od -cb junk
0000000   n   o   w       i   s       t   h   e       t   i   m   e  \n
        156 157 167 040 151 163 040 164 150 145 040 164 151 155 145 012
0000020   f   o   r       a   l   l       g   o   o   d       p   e   o
        146 157 162 040 141 154 154 040 147 157 157 144 040 160 145 157
0000040   p   l   e  \n
        160 154 145 012
0000044
$
```
The 7-digit numbers down the left side are positions in the file, that is, the ordinal number of the next character shown, in octal.
By the way, the emphasis on octal numbers is a holdover from the PDP-11, for which octal was the preferred notation.
Hexadecimal is better suited for other machines; the `-x` option tells `od` to print in hex.

Notice that there is a character after each line, with octal value `012`.
This is the ASCII *newline* character; it is what the system places in the input when you press the RETURN key.
By a convention borrowed from C, the character representation of a newline is `\n`, but this is only a convention used by programs like `od` to make it easy to read - the value stored in the file is the single byte `012`.

Newline is the most common example of a *special character*.
Other characters associated with some terminal control operation include backspace (octal value `010`, printed as `\b`), tab (`011`, `\t`), and a carriage return (`015`, `\r`).

It is important in each case to distinguish between how the character is stored in a file and how it is interpreted in various situations.
For example, when you type a backspace on your keyboard (and assuming that your erase character is backspace), the kernel interprets it to mean that you want to discard whatever character you typed previously.
Both that character and the backspace disappear, but the backspace is echoed to your terminal, where it makes the cursor move one position backwards.

If you type the sequence
```
\⌫
```
(i.e., `\` followed by a backspace), however, the kernel interprets that to mean that you want a literal backspace in your input, so the `\` is discarded and the byte `010` winds up in your file.
When the backspace is echoed on your terminal, it moves the cursor to sit on top of the `\`.
When you *print* a file that contains a backspace, the backspace is passed uninterpreted to your terminal, which again will move the cursor one position backwards.
When you use `od` to display a file that contains a backspace, it appears as a byte with value `010`, or, with the `-c` option, as `\b`.

The story for tabs is much the same: on input, a tab character is echoed to your terminal and sent to the program that is reading; on output, the tab is simply sent to the terminal for interpretation there.
There is a difference, though - you can tell the kernel that you want *it* to interpret tabs for you on output; in that case, each tab that would be printed is replaced by the right number of blanks to get to the next tab stop.
Tab stops are set a columns 9, 17, 25, etc.
The command
```
$ stty -tabs
```
causes tabs to be replaced by spaces *when printed on your terminal*.
See `stty`(1).

The treatment of RETURN is analogous.
The kernel echoes RETURN as a carriage return and a newline, but stores only the newline in the input.
On output, the newline is expanded into carriage return and newline.

The UNIX system is unusual in its approach to representing control information, particularly its use of newlines to terminate lines.
Many systems instead provide "records," one per line, each of which contains not only your data but also a count of the number of characters in the line (and no newline).
Other systems terminate each line with a carriage return *and* a newline, because that sequence is necessary for output on most terminals.
(The word "linefeed" is a synonym for newline, so this sequence is often called "CRLF", which is nearly pronounceable.)

The UNIX system does neither - there are no records, no record counts, and no bytes in any file that your programs did not put there.
A newline is expanded into a carriage return and a newline when sent to the terminal, but programs need only deal with the single newline character, because that is all they see.
For most purposes, this simple scheme is exactly what is wanted.
When a more complicated structure is needed, it can easily be build on top of this; the converse, creating simplicity from complexity, is harder to achieve.

Since the end of a line is marked by a newline character, you might expect a file to be terminated by another special character, say `\e` for "end of file."
Looking at the output of `od`, though, you will see no special character at the end of the file - it just stops.
Rather than using a special code, the system signifies the end of a file by simplify saying there is no more data in the file.
The kernel keeps track of file lengths, so a program encounters end-of-file when it has processed all the bytes in a file.

Programs retrieve the data in a file by a system call (a subroutine in the kernel) called `read`.
Each time `read` is called, it returns the next part of a file - the next line of text typed on the terminal, for example.
`read` also says how many bytes of the file were returned, so end of file is assumed when a `read` says "zero bytes are being returned."
If there were any bytes left, `read` would have returned some of them.
Actually, it makes sense not to represent end of file by a special byte value, because, as we said earlier, the meaning of the bytes depends on the interpretation of the file.
But *all* files must end, and since all files must be accessed through `read`, returning zero is an interpretation-independent way to represent the end of a file without introducing a new special character.

When a program reads from your terminal, each input line is given to the program by the kernel only when you type its newline (i.e, press RETURN).
Therefore if you make a typing mistake, you can back up and correct it if you realize the mistake before you type newline.
If you type newline before realizing the error, the line has been read by the system and you cannot correct it.

We can see how this line-at-a-time input works by using `cat`.
`cat` normally saves up or *buffers* its output to write in large chunks for efficiency, but `cat -u` "unbuffers" the output, so it is printed immediately as it is read:
```
$ cat
123
456
789
ctl-d
123
456
789
$ cat -u
123
123
456
456
789
789
ctl-d
$
```
`cat` receives each line when you press RETURN; without buffering, it prints the data as it is received.

Now try something different: type some characters and then a `ctl-d` rather than a RETURN:
```
$ cat -u
123ctl-d123
```
`cat` prints the characters out immediately.
`ctl-d` says, "immediately send the characters I have typed to the program that is reading from my terminal."
The `ctl-d` itself is not sent to the program, unlike a newline.
Now type a second `ctl-d` with no other characters:
```
$ cat -u
123ctl-d123ctl-d$
```
The shell responds with a prompt, because `cat` reads no characters, decided that meant end of file, and stopped.
`ctl-d` sends whatever you have typed to the program that is reading from the terminal.
If you haven't typed anything, the program will therefore read no characters, and that looks like the end of the file.
That is why typing `ctl-d` logs you out - the shell sees no more input.
Of course, `ctl-d` is usually used to signal an end-of-file but it is interesting that it has a more general function.

### 2.2 What's in a file

The format of a file is determined by the programs that use it; there is a wide variety of file types, perhaps because there is a wide variety of programs.
But since file types are not determined by the file system, the kernel can't tell you the type of a file: it doesn't know it.
The `file` command makes an educated guess (we'll explain how shortly):
```
$ file /bin /bin/ed /usr/src/cmd/ed.c /usr/man/man1/ed.1
/bin:	directory
/bin/ed:	pure executable
/usr/src/cmd/ed.c:	c program text
/usr/man/man1/ed.1:	roff, nroff, or eqn input text
$
```
These are four fairly typical files, all related to the editor:
- the directory in which it resides (`/bin`)
- the "binary" or runnable program itself (`/bin/ed`)
- the "source" or C statements that define the program (`/usr/src/cmd/ed.c`)
- the manual page (`/usr/man/man1/ed.1`)

To determine the types, `file` didn't pay attention to the names (although it could have), because naming conventions are just conventions, and thus not perfectly reliable.
For example, files suffixed `.c` are almost always C source, but there is nothing to prevent you from creating a `.c` file with arbitrary contents.
Instead, `file` reads the first few hundred bytes of a file and looks for clues to the file type.
(As we will show later on, files with special system properties, such as directories, can be identified by asking the system, but `file` could identify a directory by reading it.)

Sometimes the clues are obvious.
A runnable program is marked by a binary "magic number" at its beginning.
`od` with no options dumps the file in 16-bit, or 2-byte, words and makes the magic number visible:
```
$ od /bin/ed
...
$
```
The octal value `410` marks a pure executable program, one for which the executing code may be shared by several processes.
(Specific magic numbers are system dependent.)
The bit pattern represented by `410` is not ASCII text, so this value could not have been created inadvertently by a program like an editor.
But you could certainly create such a file by running a program of your own, and the system understands the convention that such files are program binaries.

For text files, the clues may be deeper in the file, so `file` looks for words like `#include` to identify C source, or lines beginning with a period to identify `nroff` or `troff` input.

You might wonder why the system doesn't track file types more carefully, so that, for example, `sort` is never given `/bin/ed` as input.
One reason is to avoid foreclosing some useful computation.
Although
```
$ sort /bin/ed
```
doesn't make much sense, there are many commands that can operate on any file at all, and there's no reason to restrict their capabilities.
`od`, `wc`, `cp`, `cmp`, `file`, and many others process files regardless of their contents.
But the formatless idea foes deeper than that.
If, say, `nroff` input were distinguished from C source, the editor would be forced to make the distinction when it created a file, and probably when it read in a file for editing again.
And it would certainly make it harder for us to typeset the C programs in Chapters 8 through 8!

Instead of creating distinctions, the UNIX system tries to efface them.
All text consists of lines terminated by newline characters, and most programs understand this simple format.
Many times while writing this book, we ran commands to create text files, processed them with commands like those listed above, and used an editor to merge them into the `troff` input for the book.
The transcripts you see on almost every page are made by commands like
```
$ od -c junk >temp
$ ed ch2.1
1534
r temp
168
...
```
`od` produces text on its standard output, which can then be used anywhere text can be used.
This uniformity is unusual; most systems have several file formats, even for text, and require negotiation by a program or a user to create a file of a particular type.
In UNIX systems there is just one kind of file, and all that is required to access a file is its name.

The lack of file formats is an advantage overall - programmers needn't worry about file types, and all the standard programs will work on any file - but there are a handful of drawbacks.
Programs that sort and search and edit really expect text as input: `grep` can't examine binary files correctly, nor can `sort` sort them, nor can any standard editor manipulate then.

There are implementation limitations with most programs that expect text as input.
We tested a number of programs on a 30,000 byte text file containing no newlines, and surprisingly few behaved properly, because most programs make unadvertised assumptions about the maximum length of a line of text (for an exception, see the BUGS section of `sort`(1)).

Non-text files definitely have their place.
For example, very large databases usually need extra address information for rapid access; this has to be binary for efficiency.
But every file format that is not text must have its own family of support programs to do things that the standard tools could perform if the format were text.
Text files may be a little less efficient in machine cycles, but this must be balanced against the cost of extra software to maintain more specialized formats.
If you design a file format, you should think carefully before choosing a non-textual representation.
(You should also think about making your programs robust in the face of long input lines.)

### 2.3 Directories and filenames

All of the files you own have unambiguous names, starting with `/usr/you`, but if the only file you have is `junk`, and you type `ls`, it doesn't print `/usr/you/junk`; the filename is printed without any prefix:
```
$ ls
junk
$
```
That is because each running program, that is, each process, has a *current directory*, and all filenames are implicitly assumed to start with the name of that directory, unless they begin directly with a slash.
Your login shell, and `ls`, therefore, have a current directory.
The command `pwd` (print working directory) identifies the current directory:
```
$ pwd
/usr/you
$
```

The current directory is an attribute of a process, not a person or a program - people have login directories, processes have current directories.
If a process creates a child process, the child inherits the current directory of its parent.
But if the child then changes to a new directory, the parent is unaffected - its current directory remains the same no matter that the child does.

The notion of having a current directory is certainly a notational convenience, because it can save a lot of typing, but its real purpose is organizational.
Related files belong together in the same directory.
`/usr` is often the top directory of the user file system.
(`user` is abbreviated to `usr` in the same spirit as `cmp`, `ls`, etc.)
`/usr/you` is your login directory, your current directory when you first log in.
`/usr/src` contains source for system programs, `/usr/src/cmd` contains source for UNIX commands, `/usr/src/cmd/sh` contains the source files for the shell, and so on.
Whenever you embark on a new project, or whenever you have a set of related files, say a set of recipes, you could create a new directory with `mkdir` and put the files there.
```
$ pwd
/usr/you
$ mkdir recipes
$ cd recipes
$ pwd
/usr/you/recipes
$ mkdir pie cookie
$ ed pie/apple
...
$ ed cookie/choc.chip
...
$
```
Notice that it is simple to refer to subdirectories.
`pie.apple` has an obvious meaning: the apple pie recipe, in directory `/usr/you/recipes/pie`.
You could instead have put the recipe in, say, `recipes/apple.pie`, rather than in a subdirectory of `recipes`, but it seems better organized to put all the pies together, too.
For example, the crust recipe could be kept in `recipes/pie/crust` rather than duplicating it in each pie recipe.

Although the file system is a powerful organizational tool, you can forget where you put a file, or even what files you've got.
The obvious solution is a command or two to rummage around directories.
The `ls` command is certainly helpful for finding files, but it doesn't look in sub-directories.
```
$ cd
$ ls
junk
recipes
$ file *
junk:	ascii text
recipes:	directory
$ ls recipes/pie
apple
crust
$
```
This piece of the file system can be shown pictorially as:
```mermaid
graph TD
A(/usr/you)
A ---C1(junk)
A ---C2(recipes)
C2---D1(pie)
C2---D2(cookie)
D2---E1(choc.chip)
D1---E2(apple)
D1---E3(crust)
```

The command `du` (disc usage) was written to tell how much disc space is consumed by the files in a directory, including all its subdirectories.
```
$ du
6         ./recipes/pie
4         ./recipes/cookie
11        ./recipes
13        .
$
```
The filenames are obvious; the numbers are the number of disc blocks - typically 512 or 1024 bytes each - of storage for each file.
The value for a directory indicates how many blocks are consumed by all the files in that directory and its subdirectories, including the directory itself.

`du` has an option `-a`, for "all", that causes it to print out all the files in a directory.
If one of those is a directory, `du` processes that as well:
```
$ du -a
2       ./recipes/pie/apple
3       ./recipes/pie/crust
6       ./recipes/pie
3       ./recipes/cookie/choc.chip
4       ./recipes/cookie
11      ./recipes
1       ./junk
13      .
$
```

The output of `du -a` can be piped through `grep` to look for specific files:
```
$ du -a | crep choc
3       ./recipes/cookie/choc.chip
$
```
Recall from chapter 1 that the name `.` is a directory entry that refers to the directory itself; it permits access to a directory without having to know the full name.
`du` looks in a directory for files; if you don't tell it which directory, it assumes `.`, the directory you are in now.
Therefore, `junk` and `./junk` are names for the same file.

Despite their fundamental properties inside the kernel, directories sit in the file system as ordinary files.
They can be read as ordinary files.
But they can't be created or written as ordinary files - to preserve its sanity and the users' files, the kernel reserves to itself all control over the contents of directories.

The time has come to look at the bytes in a directory:
```
$ od -cb .
...
$
```

See the filenames buried in there?
The directory format is a combination of binary and textual data.
A directory consists of 16-byte chunks, the last 14 bytes of which hold the filename, padded with ASCII NUL's (which have value 0) and the first two of which tell the system where the administrative information for the file resides - we'll come back
to that.
Every directory begins with two entries `.` ("dot") and `..` ("dot-dot").
```
$ cd                Home
$ cd recipes
$ pwd
/usr/you/recipes
$ cd ..; pwd        Up one level
/usr/you
$ cd ..; pwd        Up another level
/usr
$ cd ..; pwd        Up another level
/
$ cd ..; pwd        Up another level
/                   Can't go any higher
$
```
The directory `/` is called the *root* of the file system.
Every file in the system is in the root directory or one of its subdirectories, and the root is its own parent directory.

### 2.4 Permissions

Every file has a set of *permissions* associated with it, which determine who can do what with the file.
If you're so organized that you keep your love letters on the system, perhaps hierarchically arranged in a directory, you probably don't want other people to be able to read them.
You could therefore change the permissions on each letter to frustrate gossip (or only on some of the letters, to encourage it), or you might just change the permissions on the directory containing the letters, and thwart snoopers that way.

But we must warn you: there is a special user on *every* UNIX system, called the *super-user*, who can read or modify *any* file on the system.
The special login name `root` carries super-user privileges; it is used by system administrators when they do system maintenance.
There us also a command called `su` that grants super-user status if you know the `root` password.
Thus anyone who knows the super-user password can read your love letters, so don't keep sensitive material in the file system.

If you need more privacy, you can change the data in a file so that even the super-user cannot read (or at least understand) it, using the `crypt` command (`crypt`(1)).
Of course, even `crypt` isn't perfectly secure.
A super-user can change the `crypt` command itself, and there are cryptographic attacks on the `crypt` algorithm.
The former requires malfeasance and the latter takes hard work, however, so `crypt` is in practice fairly secure.

In real life, most security breaches are due to passwords that are given away or easily guessed.
Occasionally, system administrative lapses make it possible for a malicious user to gain super-user permission.
Security issues are discussed further in some of the papers cited in the bibliography at the end of this chapter.

When you log in, you type a name and then verify that you are that person by typing a password.
The name is your login identification, or *login-id*.
But the system actually recognizes you by a number, called your user-id, or *uid*.
In fact different login-id's may have the same uid, making them indistinguishable to the system, although that is relatively rare and perhaps undesirable for security reasons.
Besides a uid, you are assigned a group identification, or *group-id*, which places you in a class of users.
On many systems, all ordinary users (as opposed to those with login-id's like `root`) are placed in a single group called `other`, but your system may be different.
The file system, and therefore the UNIX system in general, determines what you can do by the permissions granted to your uid and group-id.

The file `/etc/passwd` is the *password file*; it contains all the login information about each user.
You can discover your uid and group-id, as does the system, by looking up your name in `/etc/passwd`:
```
$ grep you /etc/passwd
you:gkmbCTrJ04COM:604:1:Y.O.A.People:/usr/you:
$
```
The fields in the password file are separated by colons and are laid out like this (as seen in `passwd`(5)):
```
login-id:encrypted-password:uid:group-id:miscellany:login-directory:shell
```
The file is ordinary text, but the field definitions and separator are a convention agreed upon by the programs that use the information in the file.

The shell field is often empty, implying that you use the default shell, `/bin/sh`.
The miscellany field may contain anything; often, it has your name and address or phone number.

Note that your password appears here in the second field, but only in an encrypted form.
Anybody can read the password file (you just did), so if your password itself were there, anyone would be able to use it to masquerade as you.
When you give your password to `login`, it encrypts it and compares the result against the encrypted password in `/etc/passwd`.
If they agree, it lets you log in.
The mechanism works because the encryption algorithm has the property that it's easy to go from the clear form to the encrypted form, but very hard to go backwards.
For example, if your password is `ka-boom`, it might be encrypted as `gkmbCTrJ04COM`, but given the latter, there's no easy way to get back to the original.

The kernel decided that you should be allowed to read `/etc/passwd` by looking at the permissions associated with the file.
There are three kinds of permissions for each file:
- read (examine its contents)
- write (change its contents)
- execute (run it as a program)
Furthermore, different permissions can apply to different people.
As file owner, you have one set of read, write, and execute permissions.
Your "group" has a separate set.
Everyone else has a third set.

The `-l` option of `ls` prints the permissions information, among other things:
```
$ ls -l /etc/passwd
-rw-r--r-- 1 root       5115 Aug 30 10:40 /etc/passwd
$ ls -lg /etc/passwd
-rw-r--r-- 1 adm        5115 Aug 30 10:40 /etc/passwd
$
```
These two lines may be collectively interpreted as: `/etc/passwd` is owned by login-id `root`, group `adm`, is 5115 bytes long, was last modified on August 30 at 10:40 AM, and has one link (one name in the file system; we'll discuss links in the next section).
Some versions of `ls` give both owner and group in one invocation.

The string `-rw-r--r--` is how `ls` represents the permissions on the file.
The first `-` indicates that it is an ordinary file.
If it were a directory, there would be `d` there.
The next three characters encode the file owner's (based on uid) read, write, and execute permissions.
`rw-` means that `root` (the owner) may read or write, but not execute the file.
An executable file would have an `x` instead of a dash.

The next three characters (`r--`) encode group permissions, in this case that people in group `adm`, presumably the system administrators, can read the file but not write or execute it.
The next three (also `r--`) define the permissions for everyone else - the rest of the users on the system.
On this machine, then, only `root` can change the login information for a user, but anybody may read the file to discover the information.
A plausible alternative would be for group `adm` to also have write permission on `/etc/passwd`.

The file `/etc/group` encodes group names and group-id's, and defines which users are in which groups.
`/etc/passwd` identifies only your login group; the `newgrp` command changes your group permissions to another group.

Anybody can say
```
$ ed /etc/passwd
```
and edit the password file, but only `root` can write back the changes.
You might therefore wonder how you can change your password, since it involves editing the password file.
The program to change passwords is called `passwd`; you will probably find it in `/bin`:
```
$ ls -l /bin/passwd
-rwsr-xr-x	1	root	8454	Jan	4	1983	/bin/passwd
$
```
(Note that `/etc/passwd` is the text file containing the login information, while `/bin/passwd`, in a different directory, is a file containing an executable program that lets you change the password information.)
The permissions state that anyone may execute the command, but only `root` can change the `passwd` command.
But the `s` instead of an `x` in the execute field for the file owner states that, when the command is run, it is to be given the permissions corresponding to the file owner, in this case `root`.
Because `/bin/passwd` is "set-uid" to `root`, any user can run the `passwd` command to edit the password file.

The set-uid bit is a simple but elegant idea that solves a number of security problems.
For example, the author of a game program can make the program set-uid to the owner, so that it can update a score file that is otherwise protected from other users' access.
But the set-uid concept is potentially dangerous.
`/bin/passwd` has to be correct; if it were not, it could destroy the system information under `root`'s' auspices.
If it had the permissions `-rwsrwxrwx`, it could be overwritten by *any* user, who could therefore replace the file with a program that does anything.
This is particularly serious for a set-uid program, because `root` has access permissions to every file on the system.
(Some UNIX systems turn the set-uid bit off whenever a file is modified, to reduce the danger of a security hole.)

The set-uid bit is powerful, but used primarily for a few system programs such as `passwd`.
Let's look at a more ordinary file.
```
$ ls -l /bin/who
-rwxrwxr-x	1	root	6348	Mar	29	1983	/bin/who
$
```
`who` is executable by everybody, and writable by `root` and the owner's group.
What executable "means" is this: when you type
```
$ who
```
to the shell, it looks in a set of directories, one of which is `/bin`, for a file named `who`.
If it finds such a file, and the file has execute permission, the shell calls the kernel to run it.
The kernel checks permissions, and, if they are valid, runs the program.
Note that a program is just a file with execute permission.
In the next chapter, we will show you programs that are just text files, but that can be executed as commands because they have execute permission set.

Directory permissions operate a little differently, but the basic idea is the same.
```
$ ls -ld .
drwxrwxr-x	3	you	80	Sep	27	06:11	.
$
```
The `-d` option of `ls` asks it to tell you about the directory itself, rather than its contents, and the leading `d` in the output signifies that `.` is indeed a directory.
An `r` field means that you can read the directory, so you can find out what files are in it with `ls` (or `od` for that matter).
A `w` means that you can create and delete files in this directory, because that requires modifying and therefore writing a directory file.

Actually, you cannot simply write in a directory -  even `root` is forbidden to do so.
```
$ who > .               Try to overwrite `.`
.: cannot create        You can't
$
```
Instead there are system calls that create and remove files, and only through them is it possible to change the contents of a directory.
The permissions idea, however, still applies: the `w` fields tell who can use the system routines to modify the directory.

Permission to remove a file is independent of the file itself.
If you have write permission in a directory, you may remove files there, even files that are protected against writing.
The `rm` command asks for confirmation before removing a protected file, however, to check that you really want to do so - one of the rare occasions that a UNIX program double-checks your intentions.
(The `-f` flag to `rm` forces it to remove files without question.)

The `x` field in the permissions on a directory does not mean execution; it means "search."
Execute permissions on a directory determines whether the directory may be searched for a file.
It is therefore possible to create a directory with the mode `--x` for other users, implying that users may access any file that they know about in that directory, but may not run `ls` on it or read it to see what files are there.
Similarly, with directory permissions `r--`, users can see (`ls`) but not use the contents of a directory.
Some installations use this device to turn off `/usr/games` during busy hours.

The `chmod` (change mode) command changes permissions on files.
```
$ chmod	permissions	filenames	...
```
The syntax of the *permissions* is clumsy however.
They can be specified in two ways, wither as octal numbers or by symbolic description.
The octal numbers are easier to use, although the symbolic descriptions are sometimes convenient because they can specify relative changes in permissions.
It would be nice if you could say
```
$ chmod rw-rw-rw- junk        Doesn't work this way!
```
rather than
```
$ chmod 666 junk
```
but you cannot.
The octal modes are specified by adding together:
- `4` for read
- `2` for write
- `1` for execute
The three digits specify, as in `ls`, permissions for the owner, group, and everyone else.
The symbolic codes are difficult to explain; you must look in `chmod`(1) for a proper description.
For our purposes, it is sufficient to note that `+` turns on a permission and that `-` turns if off.
For example
```
$ chmod +x command
```
allows everyone to execute `command`, and
```
$ chmod -w file
```
turns off write permission for everyone, including the file's owner.
Except for the usual disclaimer about super-users, only the owner of a file may change the permissions on a file, regardless of the permissions themselves.
Even if somebody else allows you to write a file, the system will not allow you to change its permissions bits.
```
$ ls -ld /usr/mary
drwxrwxrwx	5	mary	704	Sep	25	10:18	/usr/mary
$ chmod 444 /usr/mary
chmod: can't change /usr/mary
$
```
If a directory is writable, however, people can remove files in it regardless of the permissions on the files themselves.
If you want to make sure that you or your friends never delete files from a directory, remove write permission from it:
```
$ cd
$ date > temp
$ chmod -w .
$ ls -ld .
dr-xr-xr-x	3	you	80	Sep	27	11:48	.
$ rm temp
rm: temp not removed
$ chmod 775 .
$ ls -ld .
dr-xr-xr-x	3	you	80	Sep	27	11:48	.
$ rm temp
$
```
`temp` is now gone.
Notice that changing the permissions on the directory didn't change it's modification date.
The modification date reflects changes to the file's contents, not its modes.
The permissions and dates are not stored in the file itself, but in a system structure called an index node, or *inode*, the subject of the next section.

### 2.5 Inodes

A file has several components: a name, contents, and administrative information such as permissions and modification times.
The administrative information is stored in the *inode* (over the years, the hyphen fell out of "i-node"), along with essential system data such as how long it is, where on the disc the contents of the file are stored, and so on.

There are three times in the inode: the time that the contents of the file were last modified (written); the time that the file was last used (read or executed); and the time that the inode itself was last changed, for example to set permissions.
```
$ date
Tue	Sep	27	12:07:24	EDT	1983
$ date >junk
$ ls -l junk
-rw-rw-rw-	1	you			29 Sep 27 12:07 junk
$ ls -lu junk
-rw-rw-rw-	1	you			29 Sep 27 06:11 junk
$ ls -lc junk
-rw-rw-rw-	1	you			29 Sep 27 12:07 junk
$
```
Changing the contents of a file does not affect it's usage time, as reported by `ls -lu`, and changing the permissions affects only the inode change time, as reported by `ls -lc`.
```
$ chmod 444 junk
$ ls -lu junk
-rw-rw-rw-	1	you			29 Sep 27 06:11 junk
$ ls -lc junk
-rw-rw-rw-	1	you			29 Sep 27 12:11 junk
$ chmod 666 junk
$
```

The `-t` option to `ls`, which sorts the files according to time, by default that of last modification, can be combined with `-c` or `-u` to report the order in which the inodes were changed or files were read:
```
$ ls recipes
cookie
pie
$ ls -lut
drwxrwxrwx	4	you			64 Sep 27 12:11 recipes
-rw-rw-rw-	1	you			29 Sep 27 06:11 junk
$
```
`recipes` is most recently used, because we just looked at its contents.

It is important to understand inodes, not only to appreciate the options on `ls`, but because in a strong sense the inodes *are* the files.
All the directory hierarchy does is provide convenient names for files.
The system's internal name for a file is it's *i-number*: the number of the inode holding the file's information.
`ls -i` reports the i-number in decimal:
```
$ date >x
$ ls -i
15768 junk
15274 recipes
15852 x
$
```
It is the i-number that is stored in the first two bytes of a directory, before the name.
`od -d` will dump the data in decimal by byte pairs rather than octal by bytes and thus make the i-number visible.
```
TODO: 
```
The first two bytes in each directory entry are the only connection between the name of a file and its contents.
A filename in a directory is therefore called a *link*, because it links a name in the directory hierarchy to the inode, and hence to the data.
The same i-number can appear in more than one directory.
The `rm` command does not actually remove inodes; it removes directory entries or links.
Only when the last link to a file disappears does the system remove the inode, and hence the file itself.

If the i-number in a directory entry is zero, it means a link has been removed, but not necessarily the contents of the file - there may still be a link somewhere else.
You can verify that the i-number goes to zero by removing the file:
```
TODO:
```
The next file created in this directory will go into the unused slot, although it will probably have a different i-number.

The `ln` command makes a link to an existing file, with the syntax
```
$ ln old-file new-file
```
The purpose of a link is to give two names to the same file, often so it can appear in two different directories.
On many systems there is a link to `/bin/ed` called `/bin/e`, so that people can call the editor `e`.
Two links to a file point to the same inode, and hence have the same i-number:
```
$ ln junk linktojunk
$ ls -li
total 3
15768 -rw-rw-rw-	2	you			29 Sep 27 06:11 junk
15768 -rw-rw-rw-	2	you			29 Sep 27 06:11 linktojunk
15274 drwxrwxrwx	4	you			64 Sep 27 12:11 recipes
$
```
The integer printed between the permissions and the owner is the number of links to the file.
Because each link just points to the inode, each link is equally important - there is no difference between the first link and subsequent ones.
(Notice that the total disc space computed by `ls` is wrong because of double counting.)

When you change a file, access to the file by any of its names will reveal the changes, since all the links point to the same file.
```
$ echo x >junk
$ ls -l
total 3
-rw-rw-rw-	2	you			29 Sep 27 06:11 junk
-rw-rw-rw-	2	you			29 Sep 27 06:11 linktojunk
drwxrwxrwx	4	you			64 Sep 27 12:11 recipes
$ rm linktojunk
$ ls -l
-rw-rw-rw-	1	you			29 Sep 27 06:11 junk
drwxrwxrwx	4	you			64 Sep 27 12:11 recipes
$
```
After `linktojunk` is removed, the link count goes back to one.
As we said before, `rm`'ing a file just breaks a link; the file remains until the last link is removed.
In practice, of course, most files only have one link, but again we see a simple idea providing great flexibility.

A word to the hasty: once the last link to a file is gone, the data is irretrievable.
Deleted files go in the incinerator, rather than the waste basket, and there is no way to call them back from the ashes.
(There is a faint hope of resurrection. Most UNIX systems have a formal backup procedure that periodically copies changed files to some safe place like magnetic tape, from which they can be retrieved. For your own protection and peace of mind, you should know how much backup is provided on your system. If there is none, watch out - some mishap to the discs could be a catastrophe.)

Links to files are handy when two people wish to share a file, but sometimes you really want a *separate copy* - a different file with the same information.
You might copy a document before making extensive changes to it, for example, so you can restore the original if you decide you don't like changes.
Making a link wouldn't help, because when the data changed, both links would reflect the change.
`cp` makes copies of files:
```
$ cp junk copyofjunk
$ ls -li
total 3
15850 -rw-rw-rw-	1	you			29 Sep 27 06:11 copyojunk
15768 -rw-rw-rw-	1	you			29 Sep 27 06:11 junk
15274 drwxrwxrwx	4	you			64 Sep 27 12:11 recipes
$
```
The i-numbers of `junk` and `copyofjunk` are different, because they are different files, even though they currently have the same contents.
It's often a good idea to change the permissions on a backup copy so it's harder to remove accidentally.
```
TODO:
```
Changing the copy of a file doesn't change the original, and removing the copy has no effect on the original.
Notice that because `copyofjunk` had write permission turned off, `rm` asked for confirmation before removing the file.

There is one more common command for manipulating files `mv` moves or renames files, simply by rearranging the links.
It's syntax is the same as `cp` and `ln`:
```
$ mv junk sameoldjunk
$ ls -li
total 2
15274 drwxrwxrwx	4	you			64 Sep 27 09:34 recipes
15768 -rw-rw-rw-	1	you			29 Sep 27 13:16 sameoldjunk
$
```
`sameoldjunk` is the same file as our old `junk`, right down to the i-number; only its name - the directory entry associated with inode 15768 - has been changed.

We have been doing all this file shuffling in one directory, but it also works across directories.
`ln` is often used to put links with the same name in several directories, such as when several people are working on one program or document.
`mv` can move a file or directory from one directory to another.
In fact, these are common enough idioms that `mv` and `cp` have special syntax for them:
```
$ mv (or cp) file1 file2... directory
```
moves (or copies) one or more files to the directory which is the last argument.
The links or copies are made with the same filenames.
For example, if you wanted to try your hand at beefing up the editor, you might begin by saying
```
$ cp /usr/src/cmd/ed.c .
```
to get your own copy of the source to play with.
If you were going to work on the shell, which is in a number of different source files, you would say
```
$ mkdir sh
$ cp /usr/src/cmd/sh/* sh
```
and `cp` would duplicate all of the shell's source files in your subdirectory `sh` (assuming no subdirectory structure in `/usr/src/cmd/sh` - `cp` is not very clever).
On some systems, `ln` also accepts multiple file arguments, again with a directory as the last argument.
And on some systems, `mv`, `cp`, and `ln` are themselves links to a single file that examines its name to see what service to perform.

### 2.6 The directory hierarchy

In Chapter 1, we looked at the file system hierarchy rather informally, starting from `/usr/you`.
We're now going to investigate it in a more orderly way, starting from the top of the tree, the root.
The top directory is `/`.
```
$ ls /
bin
boot
dev
etc
lib
tmp
unix
usr
$ 
```
`/unix` is the program for the UNIX kernel itself: when the system starts, `/unix` is read from disk into memory and started.
Actually, the process occurs in two steps: first the file `/boot` is read; it then reads in `/unix`.
More information about this "bootstrap" process may be found in `boot`(8).
The rest of the files in `/`, at least here, are directories, each a somewhat self-contained section of the total file system.
In the following brief tour of the hierarchy, play along with the text: explore a bit in the directories mentioned.
The more familiar you are with the layout of the file system, the more effectively you will be able to use it.
Table 2.1 suggests good places to look, although some of the names are system dependent.

| directory          | description                                        |
|--------------------|----------------------------------------------------|
| `/`                | root of the file system                            |
| `/bin`             | essential programs in executable form ("binaries") |
| `/dev`             | device files                                       |
| `/etc`             | system miscellany                                  |
| `/etc/motd`        | login message of the day                           |
| `/etc/passwd`      | password file                                      |
| `/lib`             | essential libraries, etc.                          |
| `/tmp`             | temporary files; cleaned when system is restarted  |
| `/unix`            | executable form of the operating system            |
| `/usr`             | user file system                                   |
| `/usr/adm`         | system administration: accounting info., etc.      |
| `/usr/bin`         | user binaries: `troff`, etc.                       |
| `/usr/dict`        | dictionary (`words`) and support for `spell`(1)    |
| `/usr/games`       | game programs                                      |
| `/usr/include`     | header files for C programs, e.g. `math.h`         |
| `/usr/include/sys` | system header files for C programs, e.g. `inode.h` |
| `/usr/lib`         | libraries for C, FORTRAN, etc.                     |
| `/usr/man`         | on-line manual                                     |
| `/usr/man/man1`    | manual pages for section 1 of manual               |
| `/usr/mdec`        | hardware diagnostics, bootstrap programs, etc.     |
| `/usr/news`        | community service messages                         |
| `/usr/pub`         | public oddments: see `ascii`(7) and `eqnchar`(7)   |
| `/usr/src`         | source code for utilities and binaries             |
| `/usr/src/cmd`     | source for commands in `/bin` and `/usr/bin`       |
| `/usr/src/lib`     | source code for subroutine libraries               |
| `/usr/spool`       | working directories for communications programs    |
| `/usr/spool/lpd`   | line printer temporary directory                   |
| `/usr/spool/mail`  | mail in-boxes                                      |
| `/usr/spool/uucp`  | working directory for the `uucp` programs          |
| `/usr/sys`         | source code for the operating system kernel        |
| `/usr/tmp`         | alternate temporary directory (little used)        |
| `/usr/you`         | your login directory                               |
| `/usr/you/bin`     | your personal programs                             |


`/bin` (binaries) we have seen before: it is the directory where the basic programs such as `who` and `ed` reside.

`/dev` (devices) we will discuss in the next section

`etc` (*et cetera*) we have also seen before.
It contains various administrative files such as the password file and some system programs such as `/etc/gettty` which initializes a terminal connection for `/bin/login`.
`/etc/rc` is a file of shell commands that is executed after the system is bootstrapped.
`/etc/group` lists the members of each group.

`/lib` (library) contains primarily parts of the C compiler, such as `/lib/cpp`, the C preprocessor, and `/lib/libc.a`, the C subroutine library.

`/tmp` (temporaries) is a repository for short-lived files created during the execution of a program.
When you start up the editor `ed`, for example, it creates a file with a name like `/tmp/e00512` to hold its copy of the file you are editing, rather than working with the original file.
It could, of course, create the file in your current directory, but there are advantages to placing it in `/tmp`: although it is unlikely, you might already have a file called `e00512` in your directory; `/tmp` is cleaned automatically when the system starts so your directory doesn't get an unwanted file if the system crashes; and often `/tmp` is arranged on the disc for fast access.

There is a problem, of course, when several programs create files in `/tmp` at once: they might interfere with each other's files.
That is why `ed`'s temporary file has a peculiar name: it is constructed in such a way as to guarantee that no other program will choose the same name for its temporary file.
In Chapters 5 and 6 we will see ways to do this.

`/usr` is called the "user file system," although it may have little to do with the actual users of the system.
On our machine, our login directories are `/usr/bwk` and `/usr/rob`, but on your machine the `/usr` part might be different, as explained in Chapter 1.
Whether or not your personal files are in a subdirectory of `/usr`, there are a number of things you are likely to find there (although local customs may vary in this regard too).
Just as in `/`, there are directories called `/usr/bin`, `/usr/lib`, and `/usr/tmp`.
These directories have functions similar to their namesakes in `/`, but contain programs less critical to the system.
For example, `nroff` is usually in `/usr/bin` rather than `/bin`, and the FORTRAN compiler libraries live in `/usr/lib`.
Of course, just what is deemed "critical" varies from system to system.
Some systems, such as the distributed 7th Edition, have all the programs in `/bin` and do away with `/usr/bin` altogether; others split `/usr/bin` into two directories according to frequency of use.

Other directories in `/usr` are `/usr/adm`, containing accounting information and `/usr/dict`, which holds a modest dictionary (see `spell`(1)).
The on-line manual is kept in `/usr/man` - see `/usr/man/man1/spell.1` for example.
If your system has source code on-line, you will probably find it in `/usr/src`.

It is worth spending a little time exploring the file system, especially `/usr`, to develop a feeling for how the file system is organized and where you might expect to find things.

### 2.7 Devices

We skipped over `/dev` in our tour, because the files there provide a nice review of files in general.
As you might guess from the name, `/dev` contains device files.

One of the prettiest ideas in the UNIX system is the way it deals with *peripherals* - discs, tape drives, line printers, terminals, etc.
Rather than having special system routines to, for example, read magnetic tape, there is a file called `/dev/mt0` (again, local customs may vary).
Inside the kernel, references to that file are converted into hardware commands to access the tape, so if a program reads `/dev/mt0`, the contents of the tape mounted on the drive are returned.
For example,
```
$ cp /dev/mt0 junk
```
copies the contents of the tape to a file called `junk`.
`cp` has no idea there is anything special about `/dev/mt0`; it is just is file - a sequence of bytes.

The device files are something of a zoo, each creature a little different, but the basic ideas of the file system apply to each.
Here is a significantly shortened list of our `/dev`
```
$ ls -l /dev
crw--w--w-	1	root     0,   0 Sep	27	23:09	console
crw-r--r--	1	root     3,   1 Sep	27	23:09	kmem
crw-r--r--	1	root     3,   0 Sep	27	23:09	mem
brw-rw-rw-	1	root     1,  64 Sep	27	23:09	mt0
crw-rw-rw-	1	root     3,   2 Sep	27	23:09	null
crw-rw-rw-	1	root     4,  64 Sep	27	23:09	rmt0
brw-r-----	1	root     2,   0 Sep	27	23:09	rp00
brw-r-----	1	root     2,   1 Sep	27	23:09	rp01
crw-r-----	1	root     13,  0 Sep	27	23:09	rrp00
crw-r-----	1	root     13,  1 Sep	27	23:09	rrp01
crw-rw-rw-	1	root     2,   0 Sep	27	23:09	tty
crw--w--w-	1	you      1,   0 Sep	27	23:09	tty0
crw--w--w-	1	root     1,   1 Sep	27	23:09	tty1
crw--w--w-	1	root     1,   2 Sep	27	23:09	tty2
crw--w--w-	1	root     1,   3 Sep	27	23:09	tty3
$
```
The first things to notice are that instead of a byte count there is a pair of small integers, and that the first character of the mode is always a `b` or a `c`.
This is how `ls` prints the information from an inode that specifies a device rather than a regular file.
The inode of a regular file contains a list of disc blocks that store the file's contents.
For a device file, the inode instead contains the internal name for the device, which consists of its type - *character* (`c`) or *bloc* (`b`) - and a pair of numbers, called the *major* and *minor* device numbers.
Discs and tapes are block devices; everything else - terminals, printers, phone lines, etc. - is a character device.
The major number encodes the type of device, while the minor number distinguishes different instances of the device.
For example, `/dev/tty0` and `/dev/tty1` are two ports on the same terminal controller, so they have the same major device number but different minor numbers.

Disc files are usually named after the particular hardware variant they represent.
`/dev/rp00` and `/dev/rp01` are named after the DEV RP06 disc drive attached to the system.
There is just one drive, divided logically into two file systems.
If there were a second drive, its associated files would be named `/dev/rp10` and `/dev/rp11`.
The first digit specifies the physical drive, and the second which portion of the drive.

You might wonder why there are several disc device files, instead of just one.
For historical reasons and for ease of maintenance, the file system is divided into smaller subsystems.
The files in a subsystem are accessible through a directory in the main system.
The program `/etc/mount` reports the correspondence between device files and directories:
```
$ /etc/mount
rp01 on /usr
$
```
In our case, the root system occupies `/dev/rp00` (although this isn't reported by `/etc/mount`) while the user file system - the files in `/usr` and its subdirectories - reside on `/dev/rp01`.

The root file system has to be present for the system to execute.
`/bin`, `/dev`, and `/etc` are always kept on the root system, because when the system starts only files in the root system are accessible, and some files such as `/bin/sh` are needed to run at all.
During the bootstrap operation, all the file systems are checked for self-consistency (see `icheck`(8) or `fsck`(8)), and attached to the root system.
The attachment operation is called mounting, the software equivalent of mounting a new disc pack in the drive; it can normally be done only by the super-user.
After `/dev/rp01` has been mounted as `/usr`, the files in the user file system are accessible exactly as if they were part of the root system.

For the average user, the details of which file subsystem is mounted where are of little interest, but there are a couple of relevant points.
First, because the subsystem may be mounted and dismounted, it is illegal to make a link to a file in another subsystem.
For example, it is impossible to link programs in `/bin` to convenient names in private `bin` directories, because `/usr` is in a different file subsystem from `/bin`:
```
$ ln /bin/mail /usr/you/bin/m
ln: Cross-device link
$
```
There would also be a problem because inode numbers are not quite unique in different file systems.

Second, each subsystem has fixed upper limits on size (number of blocks available for files) and inodes.
If a subsystem fills up, it will be impossible to enlarge files in that subsystem until some space is reclaimed.
The `df` (disc free space) command reports the available space on the mounted file systems:
```
$ df
/dev/rp00 1989
/dev/rp01 21257
$
```
`/usr` has 21257 free blocks.
Where this is ample space or a crisis depends on how the system is used; some installations need more file space headroom than others.
By the way, of all of the commands, `df` probably has the widest variation in output format.
Your `df` output may look quite different.

Let's turn now to some more generally useful things.
When you log in, you get a terminal line and therefore a file in `/dev` through which the characters you type and receive are sent.
The `tty` command tells you which terminal you are using:
```
$ whoami
you      tty0      Sep 28 01:02
$ tty
/dev/tty0
$ ls -l /dev/tty0
crw--w--w- 1 you      1, 12 Sep 28 02:40 /dev/tty0
$ date > dev/tty0
Wed Sep 28 02:40:51 EDT 1983
$
```
Notice that you own the device, and that only you are permitted to read it.
In other words, no one else can directly read the characters you are typing.
Anyone may write on your terminal, however.
To prevent this, you could `chmod` the device, thereby preventing people from using `write` to contact you, or you could just use `mesg`.
```
$ mesg n            Turn off messages
$ ls -l /dev/tty0
crw------- 1 you      1, 12 Sep 28 02:40 /dev/tty0
$ mesg y            Restore
```

It's often useful to be able to refer by name to the terminal you are using, but it's inconvenient to determine which one it is.
The device `/dev/tty` is a synonym for your login terminal, whatever terminal you are actually using.
```
$ date > /dev/tty
Wed Sep 28 02:42:23 EDT 1983
$
```
`/dev/tty` is particularly useful when a program needs to interact with a user even though its standard input and output are connected to files rather than the terminal.
`crypt` is one program that uses `/dev/tty`.
The "clear text" comes from the standard input, and the encrypted data goes to the standard output, so `crypt` reads the encryption key from `/dev/tty`.
```
$ crypt < cleartext > cryptedtext
Enter key:                          Type encryption key
$
```
The use of `/dev/tty` isn't explicit in this example, but it is there.
If `crypt` read the key from the standard input, it would read the first line of the clear text.
So instead `crypt` opens `/dev/tty`, turns off automatic character echoing so your encryption key doesn't appear on the screen, and reads the key.
In Chapters 5 and 6 we will come across several other uses of `/dev/tty`.

Occasionally you want to run a program but don't care what output is produced.
For example, you may have already seen today's news, and don't want to read it again.
Redirecting `news` to the file `/dev/null` causes its output to be thrown away:
```
$ news > /dev/null
$
```
Data written to `/dev/null` is discarded without comment, while programs that read from `/dev/null` get end-of-file immediately, because reads from `/dev/null` always return zero bytes.

One common use of `/dev/null` is to throw away regular output so that diagnostic messages are visible.
For example, the `time` command (`time`(1)) reports the CUP usage of a program.
The information is printed on the standard error, so you can time commands that generate copious output by sending the standard output to `/dev/null`:
```
$ ls -l /usr/dict/words
-r--r--r-- 1 bin    196513 Jan 20  1979 /usr/dict/words
$ time grep e /usr/dict/words > /dev/null

real       13.0
user        9.0
sys         2.7
$ time egrep e /usr/dict/words > /dev/null

real        8.0
user        3.9
sys         2.6
$
```
The numbers in the output of `time` are elapsed clock time, CPU time spent in the program and CPU time spent in the kernel while the program was running.
`egrep` is a high-powered variant of `grep` that we will discuss in Chapter 4; it's about twice as fast as `grep` when searching through large files.
If output from `grep` and `egrep` had not been sent to `/dev/null` or a real file, we would have had to wait for hundreds of thousands of characters to appear on the terminal before finding out the timing information we were after.

