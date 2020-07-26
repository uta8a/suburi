import React from "react";

const Title = () => {
  return <h1 className="text-center pt-6 text-white text-3xl">Notification</h1>;
};
const Card = () => {
  return (
    <div className="bg-orange-200 border-l-4 border-orange-500 text-orange-700 text-xl p-4 m-2 w-4/5">
      <p className="font-bold text-2xl">title</p>
      <div className="flex flex-row justify-between">
        <p>Notification</p>
        <p>0:00:00</p>
      </div>
    </div>
  );
};

const Notifications = () => {
  return (
    <>
      <Title />
      <div className="flex flex-col items-center pt-6">
        <Card />
        <Card />
      </div>
    </>
  );
};

export { Notifications };
