import React from 'react'
import { CTF_NAME,CTF_SUBTITLE } from '../util/global'
const Home = () => {

  return (
    <div className="flex flex-col">
      <h1 className="my-6 text-center text-white text-5xl font-bold">{CTF_NAME}</h1>
      <p className="my-6 text-white text-center"><CTF_SUBTITLE /></p>
      <div className="ctf-timer my-6 text-white text-center text-xl">0:00:00</div>
    </div>
  )
}

export { Home }
