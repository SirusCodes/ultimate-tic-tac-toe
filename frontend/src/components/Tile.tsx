import { useGameStore } from '../store/useGameStore';

type TileProps = {
  zone: number;
  index: number;
};

export function Tile({ zone, index }: TileProps) {
  const state = useGameStore((selector) => selector.state[zone][index]);
  const { makeMove } = useGameStore();

  return (
    <div
      onClick={() => state === 'playable' && makeMove(zone, index)}
      className={`border font-black text-2xl md:text-5xl min-w-2 aspect-square flex items-center justify-center ${state === 'playable' ? 'bg-green-300' : ''}`}
    >
      {state === 'x' && 'X'}
      {state === 'o' && 'O'}
    </div>
  );
}
