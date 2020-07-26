import React from "react";
import { FLAG_FORMAT } from "../util/global";
import { Title } from "./view/Title";

const Rules = () => {
  return (
    <>
      <Title>Rules</Title>
      <div className="px-10 flex flex-col items-center justify-center text-white text-2xl">
        <div>
          <ul className="px-10 list-inside flex flex-col">
            <li>- 他のプレイヤーの妨害を禁止します。</li>
            <li>- フラッグやヒントのチーム外での共有を禁止します。</li>
            <li>
              -
              DoSなどの過度な負荷のかかる攻撃や、スコアサーバ自体への攻撃を禁止します。また、ツールを使ったスキャンはDoSとみなされることがあります。
            </li>
            <li>
              - flagのフォーマットは
              <code className="bg-gray-700">{FLAG_FORMAT}</code>です。
            </li>
            <li>- 質問はdiscordで受け付けています。</li>
          </ul>
        </div>
      </div>
    </>
  );
};
export { Rules };
