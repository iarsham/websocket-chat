// @ts-ignore
import homeImage from "../assets/images/home.jpg";
import React, {useState} from "react";
import "../styles/Home.css";
import axios from "axios";

const Home: React.FC = () => {
    const [roomName, setRoomName] = useState<string>("");

    const handleCreateRoom = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await axios.post("http://localhost:8000/api/rooms/", {roomName});
        } catch (error) {
            //   @ts-ignore
            toast.error(error);
        }
    };

    return (
        <div className="home">
            <img src={homeImage} alt="Chat App"/>
            <h1>Welcome to the Chat App!</h1>
            <div className="actions">
                <input
                    type="text"
                    value={roomName}
                    placeholder="Room Name"
                    onChange={(e) => setRoomName(e.target.value)}
                />
                <button type="submit" onClick={handleCreateRoom}>
                    Create Room
                </button>
            </div>
        </div>
    );
};

export default Home;
