import React from "react";
import { Title } from "./view/Title";

const Scoreboard = () => {
  const placeStyle = "text-xl text-center text-white";
  const teamStyle = "text-xl text-center text-white";
  const scoreStyle = "text-xl text-center text-white";
  let list = [];
  let data = [
    { place: 1, team: "neko-chan", score: 12000 },
    { place: 2, team: "kame-chan", score: 11000 },
    { place: 3, team: "zou-chan", score: 10000 },
    { place: 4, team: "wan-chan", score: 9000 },
    { place: 5, team: "malony-chan", score: 8000 },
    { place: 6, team: "ramen-chan", score: 7000 },
  ];
  for (let p of data) {
    list.push(
      <tr>
        <td className={placeStyle}>{p.place}</td>
        <td className={teamStyle}>{p.team}</td>
        <td className={scoreStyle}>{p.score}</td>
      </tr>
    );
  }
  return (
    <>
      <Title>Scoreboard</Title>
      <div className="pt-5 flex justify-between px-20">
        <table className="bg-gray-900 w-full table-fixed">
          <thead>
            <tr>
              <th className="border text-xl text-white w-1/5">Place</th>
              <th className="border text-xl text-white w-3/5">Team</th>
              <th className="border text-xl text-white w-1/5">Score</th>
            </tr>
          </thead>
          <tbody>{list}</tbody>
        </table>
      </div>
    </>
  );
};

export { Scoreboard };
