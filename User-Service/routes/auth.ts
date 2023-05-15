import { Request, Response, Router } from "express";
import { User, IUser } from "../models/user";
import mongoose from "mongoose";
import bcrypt from "bcryptjs";
import Joi from "joi";
import jwt from "jsonwebtoken";

const router: Router = Router();

router.post("/login", async (req: Request, res: Response) => {
  const { error } = validate(req.body);
  if (error)  res.status(400).send(error.details[0].message);

  let user: IUser | null
  user = await User.findOne({ username: req.body.username });
  if (!user)  return res.status(400).send("Invalid user or password.");
  
  const salt: string = await bcrypt.genSalt(10);
  const hashed: string = await bcrypt.hash(req.body.password, salt);
  
  // const validPassword: boolean = await bcrypt.compare(
  //   req.body.password,
  //   user.password
  // );
  if (req.body.password !== user.password) return res.status(400).send("Invalid user or password.");

  const token: string = jwt.sign(
    {
      _id: user._id,
    },
    "axdaasdFWFQWF232SegqwQrqww",
    { expiresIn: "8h" }
  );

  return res.send(JSON.stringify(token));
});

function validate(req: Request): Joi.ValidationResult {
  const schema: Joi.ObjectSchema = Joi.object({
    username: Joi.string().min(3).max(50).required(),
    password: Joi.string().min(3).max(1024).required(),
  });

  return schema.validate(req);
}

export default router;
