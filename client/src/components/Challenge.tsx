import React from "react";
import { Title } from "./view/Title";

const Card = () => {
  const info = {
    title: "Sanity Check",
    score: 100,
    solves: 120,
    tag: "sanity",
    content: "discordにあるフラグを提出してください。",
    attachment: "file",
  };
  return (
    <div className="w-1/2 px-2 py-2 relative">
      <div className="modal-open relative bg-yellow-200 border-l-4 border-yellow-500 text-orange-900 text-xl p-4 w-full">
        <div className="">
          <h1 className="font-bold">{info.title}</h1>
        </div>
        <div className="bg-blue-600 w-1/4 absolute top-0 right-0 py-1 rounded-bl-lg">
          <div className="text-white font-bold text-center">
            {info.score} pts
          </div>
        </div>
        <div className="flex flex-row pt-2">
          <div className="text-base">{info.tag}</div>
          <div className="text-base pl-10">{info.solves} solves</div>
        </div>
      </div>
      <div className="modal hidden absolute top-0">
        <div className="">{info.title}</div>
        <div className="">{info.score} pts</div>
        <div className="">{info.solves} solves</div>
        <div className="">{info.tag}</div>
        <div className="">{info.content}</div>
        <div className="">{info.attachment}</div>
      </div>
    </div>
  );
};

const Challenge = () => {
  return (
    <>
      <Title>Challenge</Title>
      <div className="flex flex-row py-6 mx-10 flex-wrap">
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
      </div>
    </>
  );
};

export { Challenge };
