# Battlesnake

Golang [battlesnake](https://play.battlesnake.com/) server for the 2020 season.

### How does it work?

1) Generate many possible future scenarios
2) Group them by the move that "our" snake makes in _this_ turn
3) Score all of the future states based on some heuristic function
4) Pick the move which scores the highest on average

This is currently running on Google App Engine and accessible at https://snake.50w.co.