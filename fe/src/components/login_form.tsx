import React, { useState } from "react";
import {
  Box,
  Button,
  CircularProgress,
  Container,
  IconButton,
  InputAdornment,
  Snackbar,
  TextField,
  Typography,
  Alert,
  Paper,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import { login } from "../api/api"; // Hàm gọi API backend

const LoginForm = ({ onLogin }: { onLogin: (token: string) => void }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: "",
    severity: "success" as "success" | "error",
  });

  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setSnackbar({ ...snackbar, open: false });

    try {
      const token = await login(username, password); // Gọi API login
      onLogin(token);
      localStorage.setItem("token", token);
      setSnackbar({ open: true, message: "Đăng nhập thành công!", severity: "success" });
      navigate("/home");
    } catch (err: any) {
      setSnackbar({
        open: true,
        message: err.message || "Sai thông tin đăng nhập",
        severity: "error",
      });
    } finally {
      setLoading(false);
    }
  };

  const handleSSOLogin = () => {
    // Gọi tới backend để chuyển hướng đến Google
    window.location.href = "http://localhost:3000/login-sso";
  };

  return (
      <Container maxWidth="sm" sx={{ mt: 10 }}>
        <Paper elevation={3} sx={{ p: 4 }}>
          <Typography variant="h5" textAlign="center" mb={3}>
            Đăng nhập
          </Typography>
          <form onSubmit={handleSubmit}>
            <TextField
                label="Tên đăng nhập"
                fullWidth
                margin="normal"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
            />
            <TextField
                label="Mật khẩu"
                fullWidth
                margin="normal"
                type={showPassword ? "text" : "password"}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                InputProps={{
                  endAdornment: (
                      <InputAdornment position="end">
                        <IconButton
                            onClick={() => setShowPassword((prev) => !prev)}
                            edge="end"
                            aria-label="toggle password visibility"
                        >
                          {showPassword ? <VisibilityOff /> : <Visibility />}
                        </IconButton>
                      </InputAdornment>
                  ),
                }}
            />

            <Box mt={3} display="flex" justifyContent="center" flexDirection="column" gap={2}>
              <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  disabled={loading}
                  fullWidth
                  startIcon={loading ? <CircularProgress size={20} /> : null}
              >
                {loading ? "Đang xử lý..." : "Đăng nhập"}
              </Button>

              <Button variant="outlined" color="secondary" fullWidth onClick={handleSSOLogin}>
                Đăng nhập với Google
              </Button>
            </Box>
          </form>
        </Paper>

        <Snackbar
            open={snackbar.open}
            autoHideDuration={3000}
            onClose={() => setSnackbar({ ...snackbar, open: false })}
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert severity={snackbar.severity} variant="filled">
            {snackbar.message}
          </Alert>
        </Snackbar>
      </Container>
  );
};

export default LoginForm;
