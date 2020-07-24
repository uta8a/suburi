import React from 'react'
import { Link } from 'react-router-dom'

const AllRoute = () => {

  return (
    <div>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/user-check">User</Link>
        </li>
        <li>
          <Link to="/tester-check">Tester</Link>
        </li>
        <li>
          <Link to="/admin-check">Admin</Link>
        </li>
      </ul>
    </div>
  )
}

export { AllRoute }
