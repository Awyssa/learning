What is a Virtual Environment?

A virtual environment is an isolated Python environment that allows you to install packages and dependencies specific to a project without affecting your system's global Python installation. This helps avoid version conflicts between projects and keeps your system Python clean.

Setting Up a Virtual Environment
~ python -m venv myproject_env
~ source myproject_env/bin/activate


Project Structure:
Copymyproject/
│
├── myproject_env/        # Virtual environment folder
├── src/                  # Source code
├── tests/               # Test files
├── requirements.txt     # Project dependencies
└── .gitignore          # Git ignore file

# After activating virtual environment
# Install packages
pip install package_name

# Save dependencies
pip freeze > requirements.txt

# Install from requirements.txt
pip install -r requirements.txt


# .gitignore content
myproject_env/
__pycache__/
*.pyc