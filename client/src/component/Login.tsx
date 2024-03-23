import React, {useState} from "react";
import axios from "axios";
import "../styles/Login.css";
import {toast} from "react-toastify";
import {NavigateFunction, useNavigate} from "react-router-dom";

const Login: React.FC = () => {
    const [userName, setUserName] = useState<string>("");
    const [passWord, setPassWord] = useState<string>("");
    const navigate: NavigateFunction = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await axios.post(
                "http://localhost:8000/api/auth/login",
                {
                    userName,
                    passWord,
                },
                {
                    withCredentials: true,
                },
            ).then(
                () => {
                    toast.success("Login Successful!");
                    setTimeout(() => navigate("/"), 2000);
                },
                (error) => {
                    toast.error(error.response.data["response"]);
                },
            );
        } catch (error) {
            console.log(error)
        }
    };

    return (
        <div className="login">
            <h1>Log In Your Account</h1>
            <input
                type="text"
                value={userName}
                placeholder="Username"
                onChange={(e) => setUserName(e.target.value)}
            />
            <input
                type="password"
                value={passWord}
                placeholder="Password"
                onChange={(e) => setPassWord(e.target.value)}
            />
            <button type="submit" onClick={handleLogin}>
                Login
            </button>
        </div>
    );
};

export default Login;
