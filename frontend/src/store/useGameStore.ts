import { create } from 'zustand';
import { combine, persist } from 'zustand/middleware';

export type Player = 'x' | 'o';
export type GameState = (Player | 'playable' | null)[][];
export type WonZones = (Player | null)[];

export const useGameStore = create(
  persist(
    combine(
      {
        state: Array(9).fill(Array(9).fill('playable')) as GameState,
        currentPlayer: 'x' as Player,
        wonZones: Array(9).fill(null) as (Player | null)[],
      },
      (set) => ({
        makeMove: (zone: number, index: number) => {
          set((prev) => {
            const newState = [...prev.state.map((row) => [...row])];

            // Reset all playable tiles
            for (let i = 0; i < 9; i++) {
              for (let j = 0; j < 9; j++) {
                if (newState[i][j] === 'playable') {
                  newState[i][j] = null;
                }
              }
            }

            // Make the move
            newState[zone][index] = prev.currentPlayer;

            // Check if the move wins the zone
            const winner = isWon(newState, zone);
            const wonZones = [...prev.wonZones];
            if (winner) {
              wonZones[zone] = winner;
            }

            const playableZones = getPlayableZones(newState, wonZones, index);

            for (const playableZone of playableZones) {
              for (let i = 0; i < 9; i++) {
                if (newState[playableZone][i] === null) {
                  newState[playableZone][i] = 'playable';
                }
              }
            }

            const nextPlayer = prev.currentPlayer === 'x' ? 'o' : 'x';
            return { state: newState, currentPlayer: nextPlayer, wonZones };
          });
        },
        reset: () => {
          set(() => ({
            state: Array(9).fill(Array(9).fill('playable')) as GameState,
            currentPlayer: 'x' as Player,
            wonZones: Array(9).fill(null) as (Player | null)[],
          }));
        },
      })
    ),
    {
      name: 'game-storage',
    }
  )
);

function getPlayableZones(
  state: GameState,
  wonZones: (Player | null)[],
  lastMoveIndex: number
): number[] {
  if (wonZones[lastMoveIndex] || isZoneFull(state, lastMoveIndex)) {
    return Array.from({ length: 9 }, (_, i) => i).filter(
      (zone) => !wonZones[zone] && !isZoneFull(state, zone)
    );
  }

  return [lastMoveIndex];
}

function isZoneFull(state: GameState, zone: number): boolean {
  return state[zone].every((tile) => tile !== null && tile !== 'playable');
}

function isWon(state: GameState, zone: number): Player | null {
  // Check rows
  for (let i = 0; i < 3; i++) {
    if (
      state[zone][i * 3] !== 'playable' &&
      state[zone][i * 3] === state[zone][i * 3 + 1] &&
      state[zone][i * 3] === state[zone][i * 3 + 2]
    ) {
      return state[zone][i * 3] as Player;
    }
  }

  // Check columns
  for (let i = 0; i < 3; i++) {
    if (
      state[zone][i] !== 'playable' &&
      state[zone][i] === state[zone][i + 3] &&
      state[zone][i] === state[zone][i + 6]
    ) {
      return state[zone][i] as Player;
    }
  }

  // Check diagonals
  if (
    state[zone][0] !== 'playable' &&
    state[zone][0] === state[zone][4] &&
    state[zone][0] === state[zone][8]
  ) {
    return state[zone][0] as Player;
  }
  if (
    state[zone][2] !== 'playable' &&
    state[zone][2] === state[zone][4] &&
    state[zone][2] === state[zone][6]
  ) {
    return state[zone][2] as Player;
  }

  return null;
}
