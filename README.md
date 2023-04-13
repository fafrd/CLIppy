# CLIppy

    /‾‾\
    |  |
    O  O
    || |/
    || ||
    |\_/|
    \___/
    
    It looks like you're trying to run a Minecraft server.
    Can I help you with that?

Unfinished project. having trouble reading terminal output in a discrete manner (dividing output from each command into a specific segment for parsing.) I experimented with using tmux and scraping the output but it's very brittle. Best solution might be to use `script(1)` and some ugly string hacking to parse out just the last command's output, similar to what I did for [aquarium](https://github.com/fafrd/aquarium)... another option might be some kind of command wrapper
