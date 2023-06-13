import argparse
import yaml
import os

# Consoledot Starter App Fork Script
# This script is used to fork the starter app for a new project. It will ask the user a series of questions
# and then use the answers to customize the starter app for the new project.
# The goal is to automate the process of removing / updating boilerplate code from the starter app
# so that the developer can focus on the new project.

# The Mutator class is used to change a substring with a specific value in multiple files
# to a new value. The Mutator class is initialized with a yaml file that contains the
# information needed to perform the mutation such as the string to look for, the files to
# change, and the prompt to display to the user.
class Mutator:
    def __init__(self, mutation_file, dryrun=True):
        opts = self.read_mutation_yaml_file(mutation_file)
        self.new_value = ""
        self.old_value = opts["old_value"]
        self.files = opts["files"]
        self.prompt = opts["prompt"]
        self.mutation_name = opts["mutation_name"]
        self.dryrun = dryrun

    def set_dryrun(self, dryrun):
        self.dryrun = dryrun

    def print_prompt(self):
        print(self.prompt)

    def set_new_value(self, new_value):
        self.new_value = new_value

    def print_new_value_confirmation(self):
        print(f"{self.mutation_name} -> {self.new_value}")

    def read_mutation_yaml_file(self, path):
        with open(path) as f:
            opts = yaml.safe_load(f)
        return opts

    def get_new_value(self):
        print(self.prompt)
        self.new_value = input("> ")

    def mutate(self):
        print(f"🍴 Changing {self.mutation_name}")
        for file in self.files:
            self.mutate_file(file)

    def mutate_file(self, file):
        #If we are in dry run print and bail
        if self.dryrun:
            print(f"    Dry Run: Would have changed {self.old_value} to {self.new_value} in {file}")
            return
        #open the file for read and write
        with open(file, 'r+') as f:
            #read the file into memory
            data = f.read()
            #replace the old string with the new string
            data = data.replace(self.old_value, self.new_value)
            #set the file pointer to the beginning of the file
            f.seek(0)
            #write the data to the file
            f.write(data)
            #set the file pointer to the end of the file
            f.truncate()
        #close the file
        f.close()

# The Controller class is used to orchestrate the mutators. It is initialized with a list of
# mutators. It will then set the dryrun value on each mutator , get and confirm the user input,
# and then call the mutate method on each mutator.
class Controller:
    def __init__(self, mutators):
        self.init_args()
        self.mutators = mutators
        self.set_mutators_dryrun()
    
    def init_args(self):
        parser = argparse.ArgumentParser()
        parser.add_argument("-d", "--dryrun", type=bool, help="Dry run. Print out the config and exit.", default=False)
        self.args = parser.parse_args()

    def set_mutators_dryrun(self):
        for mutator in self.mutators:
            mutator.set_dryrun(self.args.dryrun)

    def set_mutator_new_values(self):
        for mutator in self.mutators:
            mutator.print_prompt()
            mutator.set_new_value(input("> "))
        if self.confirm_mutator_new_values() == False:
            self.set_mutator_new_values()

    def run_mutators(self):
        for mutator in self.mutators:
            mutator.mutate()

    def confirm_mutator_new_values(self):
        print("🍴 Here is what we are going to change:")
        for mutator in self.mutators:
            mutator.print_new_value_confirmation()
        print("Does this look good? (y/n)")
        confimation = input("> ")
        return confimation == "y"
    
def main():
    print("🍴 Welcome to the Starter App Fork Script!\n")
    print("This script will ask you a few questions and then use that to customize the starter app for you.\n")

    controller = Controller([
        Mutator("scripts/mutations/github_repo.yaml"),
        Mutator("scripts/mutations/quay_repo.yaml"),
        Mutator("scripts/mutations/contact_name.yaml"),
        Mutator("scripts/mutations/contact_email.yaml"),
        Mutator("scripts/mutations/api_path.yaml"),
        Mutator("scripts/mutations/api_doc_description.yaml"),
        Mutator("scripts/mutations/api_doc_title.yaml"),
    ])

    controller.set_mutator_new_values()
    controller.run_mutators()

    print("🍴 All done! The app is all yours now!\n")

if __name__ == "__main__":
    main()
