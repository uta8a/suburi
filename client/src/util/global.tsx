import React from "react";
const URL = "http://localhost:8080";
const CTF_NAME = "ErilyCTF";
const CTF_ADMIN_NAME = "@kaito_tateyama";
const CTF_ADMIN_LINK = "https://twitter.com/kaito_tateyama";
const CTF_SUBTITLE = () => {
  return (
    <>
      🚩🚩🚩 これは試験的に行われます。フィードバックは
      <a className="bg-gray-700 hover:bg-gray-600" href="#">
        こちら
      </a>
      からお願いします。 🚩🚩🚩
    </>
  );
};
const FLAG_FORMAT = 'Erily{THIS_IS_FLAG}'
export { URL, CTF_NAME, CTF_ADMIN_NAME, CTF_ADMIN_LINK, CTF_SUBTITLE, FLAG_FORMAT };
