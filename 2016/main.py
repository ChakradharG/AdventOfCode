import sys
import importlib

def main() -> None:
	try:
		solution = importlib.import_module(f"day{sys.argv[1]}.solution")
	except IndexError:
		print("Please provide a day to run (e.g., '01', '02').")
	except ModuleNotFoundError:
		print(f"Module day{sys.argv[1]} not found.")
	else:
		solution.run()

main()
