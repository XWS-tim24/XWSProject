import { Router } from 'express';
import bcrypt from 'bcryptjs';
import { User, IUser, validateUser } from '../models/user';
import { authorization } from '../middleware/authorization';
import { Request, Response} from 'express';
const router: Router = Router();

router.post('/register', async (req: Request, res: Response) => {
  const { error } = validateUser(req.body);
  if (error) return res.status(400).send(error.details[0].message);

  let user: IUser | null 
  user = await User.findOne({ username: req.body.username });
  if (user) return res.status(400).send('User already registered.');

  const salt: string = await bcrypt.genSalt(10);
  const hashed: string = await bcrypt.hash(req.body.password, salt);

  const role: string = req.body.userRole;
  if(role !== "Guest" && role !== "Host") return res.status(400).send('Bad user role.');

  user = new User({
    username: req.body.username,
    password: req.body.password,
    name: req.body.name,
    surname: req.body.surname,
    email: req.body.email,
    city: req.body.city,
    userRole: role
  });

  user = await user.save();

  return res.send({
    id: user._id,
    username: user.username,
    password: user.password,
    name: user.name,
    surname: user.surname,
    email: user.email,
    city: user.city,
    userRole : user.userRole
  });
});

interface AuthRequest extends Request {
  user?: any;
}

router.get('/currentUser', authorization, async (req: AuthRequest, res: Response) => {
  let user : IUser | null;
  user = await User.findById(req.user._id).select('-password -__v');
  if(!user) return res.status(400).send("Bad request!");
  
  return res.send(user);
});

router.get('/currentUserID', authorization, async (req: AuthRequest, res: Response) => {
  let user : IUser | null;
  
  user = await User.findById(req.user._id).select('-password -__v');
  if(!user) return res.status(400).send("Bad request!");
  
  return res.send(user.id);
});

router.get('/getUserByID/:id', async (req: Request, res: Response) => {
  let user : IUser | null;
  const id = req.params["id"]
  user = await User.findById(id).select('-password -__v');
  if(!user) return res.status(400).send("Bad request!");
  
  return res.send(user);
});


router.put('/updateProfile', authorization, async (req: AuthRequest, res: Response) => {
  let user : IUser | null;
  user = await User.findById(req.user._id).select('-__v');
  if(!user) return res.status(400).send("Bad request!");
  
  user.username = req.body.username;
  user.password = req.body.password;
  user.name = req.body.name;
  user.surname = req.body.surname;
  user.email = req.body.email;
  user.city = req.body.city;

  user = await user.save();

  return res.send(user);
});

export default router;
