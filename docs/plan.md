# What is the game

- 9x9 tic tac toe, is a meta game of 3x3
- It is a 3x3 game which small 3x3 games inside it
- It has total 9 - 3x3 games where each move decides where the next person can play
  - Eg. If X plays at the top left corner of center small game then O needs to play in the top left of main game
  - If the small game is full or won then the player is allowed to play anywhere else.
- You play in small game but have to lookout for the bigger game.

![How the board looks](https://en.wikipedia.org/wiki/File:Super_tic-tac-toe_rules_example.png)

# Requirements

- Game board should be as small as possible because we need to have a lot of them and RAM is expensive - Thanks to AI :)
- Store metadata about the game state (eg. who won which small game)
- Should be able to quickly find out which place can be played

# Board ideas

- Board will be a bit map as it will take the least amount of space
- We can design the board in 2 ways
  - As per the big game
    - First 9-bits will be the first row and so on
    - It will be difficult to track of small games
    - Big picture would be clear
  - As per the small games
    - First 9-bits will be the first small game and so on
    - Easier to track small games
    - Might be able to manage big picture with helper methods
- What metadata to store?
  - Which small games are filled/won
  - Next small game zone
  - Next player
  - Who Won

# Bits required

- 81x2 bits - positions
- Metadata bits
  - 9x2 bits - small game wins
  - 1x2 bit - win
  - 9 bits - next small game zone
  - 1 bit - next player

# Bit mapping

## Player board

Go only has max 64bits integer so we can depend on 64 bit integer + 32 bit integer or 2x64 bit integer

### Low Integer 

- 63bits - top, middle rows and bottom left small games

### High Integer

- 18 bits - bottom center and right small games
- 9 bits - small game wins

## Game Metadata

- 9 bits - next small game zone
- 9 bits - small game wins
- 1 bit - next player
