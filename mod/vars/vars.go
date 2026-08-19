
package vars

var (
    Off    = "\x1b[0m"    // Text Reset
    Black  = "\x1b[1;30m" // Black
    Red    = "\x1b[1;31m" // Red
    Green  = "\x1b[1;32m" // Green
    Yellow = "\x1b[1;33m" // Yellow
    Blue   = "\x1b[1;34m" // Blue
    Purple = "\x1b[1;35m" // Purple
    Cyan   = "\x1b[1;36m" // Cyan
    White  = "\x1b[1;37m" // White

    RedBase    = "\x1b[0;31m" // Red no highlighted
    Greenbase  = "\x1b[0;32m" // Green no highlighted
    YellowBase = "\x1b[0;33m" // Yellow no highlighted
    BlueBase   = "\x1b[0;34m" // Blue no highlighted
    PurpleBase = "\x1b[0;35m" // Purple no highlighted
    CyanBase   = "\x1b[0;36m" // Cyan no highlighted
    WhiteBase  = "\x1b[0;37m" // White no highlighted

    RedUnderline = "\x1b[4;31m" // Red underline
    OneLineUP    = "\x1b[A"

    ClearLine   = "\x1b[0G\x1b[2K\x1b[0m\r"
    ClearScreen = "\x1b[H\x1b[2J"
)

