import { useState } from "react";

type Board = {
  Name: string;
  ID: string;
};

const useApiEndpoints = () => {
  const [boards, setBoards] = useState<Board[]>([]);

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

  const getPlayers = async (boardId: string) => {
    const result = await fetch("http://localhost:3333/getPlayers?board=" + boardId)
    if (!result.ok) {
      console.log("Whoopsie")
    } else {
      console.log("OK!");
      const response = await result.json();
      console.log(response);
    }
  }

  return { boards, getBoards, addNewBoard, getPlayers }
}

export default useApiEndpoints