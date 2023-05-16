import mongoose from "mongoose";
import express from "express";
import usersRouter from './routes/users';
import authRouter from './routes/auth';
import cors from 'cors';

const app = express();

mongoose
  .connect("mongodb://172.17.0.1:27017/Users")
  .then(() => console.log("Connected to MongoDB..."))
  .catch((err) => console.error(`Could not connect to MongoDB... + ${err}`));

app.use(express.json());
app.use(cors());
app.use('/api/users', usersRouter);
app.use('/api/auth', authRouter);

const port = process.env.PORT || 5000;
const server = app.listen(port, () => console.log(`Listening on port ${port}...`));




