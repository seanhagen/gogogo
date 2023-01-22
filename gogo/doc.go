package gogo

/*
   Building a library to handle running a game of Go.

   1. The board is empty at the onset of the game (unless players
   agree to place a handicap).

   2. Black makes the first move, after which White and Black
   alternate.

   3. A move consists of placing one stone of one's own color on an
   empty intersection on the board.

   4. A player may pass their turn at any time.

   5. A stone or solidly connected group of stones of one color is
   captured and removed from the board when all the intersections
   directly adjacent to it are occupied by the enemy. (Capture of the
   enemy takes precedence over self-capture.)

   6. No stone may be played so as to recreate a former board
   position.

   7. Two consecutive passes end the game.

   8. A player's area consists of all the points the player has either
   occupied or surrounded.

   9. The player with more area wins.

   ----------------------------------

   Players and equipment

   Rule 1. Players: Go is a game between two players, called Black and White.

   Rule 2. Board: Go is played on a plain grid of 19 horizontal and 19 vertical lines, called a board.

   Definition.("Intersection", "Adjacent") A point on the board where a horizontal line meets a vertical line is called an intersection. Two intersections are said to be adjacent if they are connected by a horizontal or vertical line with no other intersections between them.

   Rule 3. Stones: Go is played with playing tokens known as stones. Each player has at their disposal an adequate supply (usually 180) of stones of the same color.

   Positions

   Rule 4. Positions: At any time in the game, each intersection on the board is in one and only one of the following three states: 1) empty; 2) occupied by a black stone; or 3) occupied by a white stone. A position consists of an indication of the state of each intersection.

   Definition. ("Connected") Two placed stones of the same color (or two empty intersections) are said to be connected if it is possible to draw a path from one intersection to the other by passing through adjacent intersections of the same state (empty, occupied by white, or occupied by black).

   Definition. ("Liberty") In a given position, a liberty of a stone is an empty intersection adjacent to that stone or adjacent to a stone which is connected to that stone.

   Play

   Rule 5. Initial position: At the beginning of the game, the board is empty.

   Rule 6. Turns: Black moves first. The players alternate thereafter.

   Rule 7. Moving: When it is their turn, a player may either pass (by announcing "pass" and performing no action) or play. A play consists of the following steps (performed in the prescribed order):

   Step 1. (Playing a stone) Placing a stone of their color on an empty intersection (chosen subject to Rule 8 and, if it is in effect, to Optional Rule 7A). It can never be moved to another intersection after being played.

   Step 2. (Capture) Removing from the board any stones of their opponent's color that have no liberties.

   Step 3. (Self-capture) Removing from the board any stones of their own color that have no liberties.

   Optional Rule 7A. Prohibition of suicide: A play is illegal if one or more stones of that player's color would be removed in Step 3 of that play.

   Rule 8. Prohibition of repetition: A play is illegal if it would have the effect (after all steps of the play have been completed) of creating a position that has occurred previously in the game.

   End

   Rule 9. End: The game ends when both players have passed consecutively. The final position is the position on the board at the time the players pass consecutively.

   Definition. ("Territory") In the final position, an empty intersection is said to belong to a player's territory if all stones adjacent to it or to an empty intersection connected to it are of that player's color.

   Definition. ("Area") In the final position, an intersection is said to belong to a player's area if either: 1) it belongs to that player's territory; or 2) it is occupied by a stone of that player's color.

   Definition. ("Score") A player's score is the number of intersections in their area in the final position.

   Rule 10. Winner: If one player has a higher score than the other, then that player wins. Otherwise, the game is a draw.
*/
