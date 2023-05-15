import Joi, { ValidationResult } from "joi";
import mongoose, { Document, Model, Schema } from "mongoose";

export interface IUser extends Document {
  username: string;
  password: string;
  name: string;
  surname: string;
  email: string;
  city: string;
  userRole: string;
}

const userSchema: Schema<IUser> = new mongoose.Schema({
  username: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 30,
    unique: true,
  },
  password: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 1024,
  },
  name: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 30,
  },
  surname: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 30,
  },
  email: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 30,
  },
  city: {
    type: String,
    required: true,
    minlength: 1,
    maxlength: 30,
  },
  userRole: {
    type: String,
    enum: ["Guest", "Host"],
  },
});

const User: Model<IUser> = mongoose.model<IUser>("User", userSchema);

function validateUser(user: IUser) : Joi.ValidationResult {
  const schema = Joi.object({
    username: Joi.string().min(1).max(50).required(),
    password: Joi.string().min(1).max(1024).required(),
    name: Joi.string().min(1).max(50).required(),
    surname: Joi.string().min(1).max(50).required(),
    email: Joi.string().min(1).max(50).required(),
    city: Joi.string().min(1).max(50).required(),
    userRole: Joi.string().valid("Guest", "Host")
  });

  return schema.validate(user);
}

export { User, validateUser };
