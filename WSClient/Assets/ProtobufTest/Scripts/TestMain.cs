using Google.Protobuf;
using Main;
using UnityEngine;
using UnityEngine.UI;
using WebSocketSharp;

namespace ProtobufTest
{
    public class TestMain : MonoBehaviour
    {
        public InputField serverInputField;

        public void WsPing()
        {
            var ws = new WebSocket("ws://" + serverInputField.text + "/echo");
            ws.OnOpen += (sender, e) => { Debug.Log(" connect to " + "ws://" + serverInputField.text + "/echo"); };
            ws.OnMessage += (sender, e) =>
            {
                Debug.Log("Pong: " + Deserializer(e.RawData).Code);
            };
            ws.OnError += (sender, e) => { Debug.Log(" error: " + e.Message); };
            ws.OnClose += (sender, e) => { Debug.Log(" connect closed: " + e.Reason); };
            ws.Connect();
            ws.Send(Serializer());
        }
    
        public void WsLogin()
        {
            var ws = new WebSocket("ws://" + serverInputField.text + "/login");
            ws.OnOpen += (sender, e) => { Debug.Log(" connect to " + "ws://" + serverInputField.text + "/login"); };
            ws.OnMessage += (sender, e) =>
            {
                Debug.Log("Success: " + DeserializerUser(e.RawData));
            };
            ws.OnError += (sender, e) => { Debug.Log(" error: " + e); };
            ws.OnClose += (sender, e) => { Debug.Log(" connect closed: " + e.Reason); };
            ws.Connect();
            ws.Send(SerializerLogin());
        }

        private byte[] Serializer()
        {
            var ping = new CommandPing
            {
                Code = 102
            };

            return ping.ToByteArray();
        }
    
        private CommandPong Deserializer(byte[] dataBytes)
        {
            IMessage message = new CommandPong();
            return (CommandPong)message.Descriptor.Parser.ParseFrom(dataBytes);
        }
    
        private byte[] SerializerLogin()
        {
            var request = new LoginRequest
            {
                UserName = "athekf"
            };

            return request.ToByteArray();
        }
    
        private User DeserializerUser(byte[] dataBytes)
        {
            IMessage message = new User();
            return (User)message.Descriptor.Parser.ParseFrom(dataBytes);
        }
    }
}