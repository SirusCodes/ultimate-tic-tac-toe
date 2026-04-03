import { ThreeByThree } from './components/ThreeByThree';
import { useGameStore } from './store/useGameStore';

function App() {
  const currentPlayer = useGameStore((state) => state.currentPlayer);
  const reset = useGameStore((state) => state.reset);

  return (
    <div className="flex flex-col m-2">
      <div className="w-screen h-screen flex flex-col items-center justify-center gap-6">
        <div className="flex justify-between w-[80vmin] ">
          <button
            className="bg-red-700 text-white px-4 py-2 rounded cursor-pointer"
            onClick={reset}
          >
            Restart
          </button>
          <div className="text-2xl font-bold">
            Current Player: {currentPlayer.toUpperCase()}
          </div>
        </div>

        <div className="grid grid-cols-3 grid-rows-3 w-[80vmin] aspect-square">
          {Array.from({ length: 9 }).map((_, index) => (
            <ThreeByThree key={index} zone={index} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
