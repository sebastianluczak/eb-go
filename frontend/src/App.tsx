import { useState } from 'react';
import './App.css'
import useApiEndpoints from './hooks/useApiEndpoints'

function App() {
  const { boards, getBoards, addNewBoard, getPlayers, addPlayer } = useApiEndpoints();
  const [playerName, setPlayerName] = useState<string>("");

  const onClickButton = () => {
    getBoards();
  }

  return (
    <>
      <div className="card">
        <label htmlFor="playerName" style={{ display: 'block', marginBottom: '8px', fontWeight: 'bold' }}>Input player name</label>
        <input
          type="text"
          name="playerName"
          id="playerName"
          style={{
            width: '100%',
            padding: '8px',
            borderRadius: '4px',
            border: '1px solid #ccc',
            boxSizing: 'border-box'
          }}
          onChange={(val) => setPlayerName(val.target.value)}
        />
      </div>
      <div className="card">
        <button onClick={onClickButton}>
          Refresh Boards
        </button>
        <button onClick={addNewBoard}>Add New Board</button>
        <p>
          {boards.map((b) => (
            <li key={b.ID}>
              <button onClick={() => getPlayers(b.ID)}>{b.Name}</button>
              {playerName && (
                <button onClick={() => addPlayer(b.ID, playerName)}>Join Board</button>
              )}
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

export default App;