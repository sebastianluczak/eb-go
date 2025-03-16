import './App.css'
import useApiEndpoints from './hooks/useApiEndpoints'

function App() {
  const { boards, getBoards, addNewBoard } = useApiEndpoints();

  const onClickButton = () => {
    console.log("Clicked");
    getBoards();
  }

  return (
    <>
      <div className="card">
        <button onClick={onClickButton}>
          Refresh Boards
        </button>
        <button onClick={addNewBoard}>Add New Board</button>
        <p>
          {boards.map((b: string) => (
            <li key={b}>{b}</li>
          ))}
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
