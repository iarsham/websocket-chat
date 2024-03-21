import React from "react";
import "../styles/Header.css";
import axios from "axios";
import {toast} from "react-toastify";
import {NavigateFunction, useNavigate} from "react-router-dom";

interface HeaderProps {
    isUserLoggedIn: boolean;
}

const Header: React.FC<HeaderProps> = ({isUserLoggedIn}) => {
    const navigate: NavigateFunction = useNavigate();

    const handleLogout = async () => {
        try {
            await axios.post(
                "http://localhost:8000/api/users/logout",
                null,
                {withCredentials: true}
            ).then(
                () => {
                    toast.success("Logout Successful!");
                    setTimeout(() => navigate("/home"), 2000);
                },
                (error) => {
                    toast.error(error.response.data["response"]);
                },
            );
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <header className="header">
            <h1>Websocket Chat APP</h1>
            <nav>
                {isUserLoggedIn ? (
                    <button className="button" type="button" onClick={handleLogout}>
                        Logout
                    </button>
                ) : (
                    <>
                        <a href="/register">SignUp</a>
                        <a href="/login">Login</a>
                    </>
                )}
            </nav>
        </header>
    );
};

export default Header;
