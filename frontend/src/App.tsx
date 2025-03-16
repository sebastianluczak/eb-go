import './App.css'
import useApiEndpoints from './hooks/useApiEndpoints'

function App() {
  const { boards, getBoards, addNewBoard, getPlayers } = useApiEndpoints();

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
          {boards.map((b) => (
            <li key={b.ID}>
              <button onClick={() => getPlayers(b.ID)}>{b.Name}</button>
            </li>
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
