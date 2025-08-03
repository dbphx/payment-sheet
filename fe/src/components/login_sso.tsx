import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const LoginSSOCallback = ({ onLogin }: { onLogin: (token: string) => void }) => {
    const navigate = useNavigate();

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get("token");
        const username = params.get("username");

        if (token && username) {
            localStorage.setItem("token", token);
            localStorage.setItem("username", username);
            onLogin(token); // gọi về App để setToken
            navigate("/home");
        } else {
            navigate("/login");
        }
    }, [navigate, onLogin]);

    return <p>Đăng nhập SSO thành công, đang chuyển hướng...</p>;
};

export default LoginSSOCallback;
