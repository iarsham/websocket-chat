import React, {useState} from "react";
import {NavigateFunction, useNavigate} from "react-router-dom";
import axios from "axios";
import "../styles/Register.css";
import {toast} from "react-toastify";

const Register: React.FC = () => {
    const [userName, setUserName] = useState<string>("");
    const [passWord, setPassWord] = useState<string>("");
    const navigate: NavigateFunction = useNavigate();

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await axios.post("http://localhost:8000/api/auth/register", {
                userName,
                passWord,
            });
            toast.success("Registration Successful!");
            setTimeout(() => navigate("/login"), 2000);
        } catch (error) {
            // @ts-ignore
            toast.error(error.response.data["response"]);
        }
    };

    return (
        <div className="register">
            <h1>Create Your Account</h1>
            <input
                type="text"
                value={userName}
                placeholder="Username"
                onChange={(e) => setUserName(e.target.value)}
                required
            />

            <input
                type="password"
                value={passWord}
                placeholder="Password"
                onChange={(e) => setPassWord(e.target.value)}
                required
            />

            <button type="submit" onClick={handleRegister}>
                Register
            </button>
        </div>
    );
};

export default Register;
