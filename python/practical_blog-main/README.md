This is a sample Django blog app for Practical Python
https://practical.learnpython.dev/

To run this example:

Create a virtual environment
```zsh
$ python3 -m venv env
$ source env/bin/activate
```

Install requirements
```zsh
(env) $ python -m pip install -r requirements.txt
```

Migrate the Database
```zsh
(env) $ python manage.py migrate
```

Create a super user
```zsh
(env) $ python manage.py createsuperuser
```

Run the server
```zsh
(env) $ python manage.py runserver
```

Every time we change our models, the database needs to be changed with it.
```zsh
(env) $ python manage.py makemigrations blog
(env) $ python manage.py migrate
```

We can easily work with models on the command line, with the built-in Django shell command.
```zsh
(env) $ python manage.py shell 
>>> from blog.models import Post
>>> Post.objects.all()
<QuerySet []>
>>> first_post = Post(title="First Post!", body="This is my first blog post.", slug="first-post")
>>> first_post.save()
>>> Post.objects.all()
<QuerySet [<Post: First Post!>]>
```

Your server should run at [http://localhost:8000/](http://localhost:8000/)