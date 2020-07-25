import React, { useState } from "react";
import { Request } from "../proto/user/user_pb";
import { UserClient } from "../proto/user/UserServiceClientPb";
import useSWR from "swr";

const LoginForm = () => {
  return (
    <div className="flex items-center justify-center mt-16 mx-6">
      <div className="p-6 max-w-sm w-full bg-gray-800 shadow rounded-md">
        <h3 className="text-white text-xl text-center">Login</h3>
        <form className="mt-4">
          <label className="block" htmlFor="username">
            <span className="text-white text-sm">Username</span>
            <input
              className="form-input py-2 px-3 mt-1 block w-full rounded-md bg-gray-800 border-2 border-yellow-500 text-white focus:border-yellow-200 focus:outline-none"
              id="username"
              type="text"
              placeholder="^[A-Za-z0-9][A-Za-z0-9_-]{0,50}$"
              autoComplete="off"
            />
          </label>

          <label className="block mt-3" htmlFor="password">
            <span className="text-white text-sm">Password</span>
            <input
              className="form-input mt-1 px-2 py-3 block w-full rounded-md bg-gray-800 border-2 border-yellow-500 text-white focus:border-yellow-200 focus:outline-none"
              id="password"
              type="password"
              name="password"
              required
              autoComplete="off"
              placeholder="******** (8 <= chars <= 50)"
            />
          </label>
          <div className="flex items-center justify-between mt-4">
            <div>
              {/* <label className="inline-flex items-center">
              <input type="checkbox" class="form-checkbox text-blue-500 bg-gray-800 border-gray-600" name="remember" id="remember" />
              <span class="mx-2 text-gray-200 text-sm">Remember Me</span> 
            </label> */}
            </div>
          </div>
          <div className="mt-6 flex content-center justify-center">
            <button
              type="submit"
              className="py-2 px-4 text-center bg-blue-700 rounded-md text-white text-sm hover:bg-blue-500 focus:outline-none"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

const Login = () => {
  return (
    <div className="Login">
      {/* login form */}
      <LoginForm />
    </div>
  );
};

export { Login };
