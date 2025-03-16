import { useState } from "react";

const useApiEndpoints = () => {
  const [boards, setBoards] = useState([]);

  const getBoards = async () => {
    console.log("Something");
    const result = await fetch("http://localhost:3333/getBoards");
    if (!result.ok) {
      console.log("Whoopsie");
    }
    const json = await result.json();

    setBoards(json);
  }

  const addNewBoard = async () => {
    const result = await fetch("http://localhost:3333/addBoard");
    if (!result.ok) {
      console.log("Whoopsie")
    } else {
      console.log("OK!");
      getBoards();
    }
  }

  return { boards, getBoards, addNewBoard }
}

export default useApiEndpoints