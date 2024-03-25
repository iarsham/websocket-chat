import React, {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import axios from "axios";
import "../styles/Room.css";

interface Message {
    id: number;
    text: string;
}

const Room: React.FC = () => {
    const {id} = useParams<{ id: string }>();
    const [roomFound, setRoomFound] = useState<boolean>(false);
    const [roomName, setRoomName] = useState<string>("");
    const [messages, setMessages] = useState<Message[]>([]);
    const [newMessage, setNewMessage] = useState<string>("");
    const navigate = useNavigate();

    useEffect(() => {
        axios
            .get("http://localhost:8000/api/rooms/", {
                withCredentials: true,
            })
            .then((res) => {
                const foundRoom = res.data.find(
                    (room: any) => room.id.toString() === id,
                );
                if (foundRoom) {
                    setRoomFound(true);
                    setRoomName(foundRoom.name);
                } else {
                    navigate("/404");
                }
            })
            .catch((error) => {
                console.log(error);
            });
    }, [id, navigate]);

    const handleSendMessage = () => {
    };

    return (
        <div className="chat-room-container">
            <h1 className="room-name">Welcome to {roomName} chat room. Start chatting!</h1>
            <div className="chat-room">
                {messages.map((message) => (
                    <div key={message.id} className="message">
                        {message.text}
                    </div>
                ))}
            </div>
            <div className="input-container">
                <input
                    type="text"
                    value={newMessage}
                    onChange={(e) => setNewMessage(e.target.value)}
                    className="input-field"
                    placeholder="Type your message..."
                />
                <button onClick={handleSendMessage} className="send-button">
                    Send
                </button>
            </div>
        </div>
    );
};

export default Room;
