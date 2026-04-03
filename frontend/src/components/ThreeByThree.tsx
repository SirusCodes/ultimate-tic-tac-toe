import { useGameStore } from '../store/useGameStore';
import { Tile } from './Tile';

type ThreeByThreeProps = {
  zone: number;
};

export function ThreeByThree({ zone }: ThreeByThreeProps) {
  const isWonBy = useGameStore((state) => state.wonZones[zone]); // Placeholder for win state logic

  return (
    <div className="@container grid grid-cols-3 grid-rows-3 border-2 aspect-square relative">
      {Array.from({ length: 9 }).map((_, index) => (
        <Tile key={index} zone={zone} index={index} />
      ))}
      {isWonBy && (
        <div className="absolute inset-0 flex items-center justify-center text-[50cqmin] font-bold text-red-500">
          {isWonBy.toUpperCase()}
        </div>
      )}
    </div>
  );
}
