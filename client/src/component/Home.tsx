// @ts-ignore
import homeImage from "../assets/images/home.jpg";
import React, {useEffect, useState} from "react";
import "../styles/Home.css";
import axios from "axios";
import {toast} from "react-toastify";
import {NavigateFunction, useNavigate} from "react-router-dom";

const Home: React.FC = () => {
    const [roomName, setRoomName] = useState<string>("");
    const [selectedRoom, setSelectedRoom] = useState<string>("");
    const [existingRooms, setExistingRooms] = useState<any[]>([]);
    const navigate: NavigateFunction = useNavigate();

    useEffect(() => {
        axios
            .get("http://localhost:8000/api/rooms/", {
                withCredentials: true,
            })
            .then((res) => {
                setExistingRooms(res.data);
            })
            .catch((error) => {
                console.log(error);
            });
    }, []);

    const handleCreateRoom = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await axios
                .post(
                    "http://localhost:8000/api/rooms/",
                    {
                        name: roomName,
                    },
                    {withCredentials: true},
                )
                .then(
                    () => {
                        toast.success("Room created successfully!");
                    },
                    (error) => {
                        toast.error(error.response.data["response"]);
                    },
                );
        } catch (error) {
            console.log(error);
        }
    };

    const handleJoinRoom = () => {
        const room = existingRooms.find((room) => room.name === selectedRoom);
        if (room) {
            navigate(`/room/${room.id}`);
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

                <div>
                    <h2>Choose from existing rooms:</h2>
                    <select
                        value={selectedRoom}
                        onChange={(e) => setSelectedRoom(e.target.value)}
                    >
                        <option value="">Rooms</option>
                        {existingRooms.map((room) => (
                            <option key={room.id} value={room.name}>
                                {room.name}
                            </option>
                        ))}
                    </select>
                    <button
                        type="submit"
                        onClick={handleJoinRoom}
                        disabled={!selectedRoom}
                        className="join"
                    >
                        Join Room
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Home;
