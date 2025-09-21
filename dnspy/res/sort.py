import shutil
from datetime import datetime


shutil.copyfile(
    "domains.txt", f"domains-{datetime.now().strftime('%Y%m%d%H%M%S')}.bak.txt"
)
with open("domains.txt", "r", encoding="utf-8") as f:
    lines = f.readlines()

lines_set = set(lines)
lines = list(lines_set)
lines.sort()

with open("domains.txt", "w", encoding="utf-8") as f:
    f.writelines(lines)


shutil.copy(
    "providers.txt", f"providers-{datetime.now().strftime('%Y%m%d%H%M%S')}.bak.txt"
)
with open("providers.txt", "r", encoding="utf-8") as f:
    lines = f.readlines()

all_hosts = list()


class HostList:
    """HostList represents a list of hosts under a specific category."""

    def __init__(self, name: str):
        self.name = name
        self.hosts = list[str]()


host_datasets = list[HostList]()
if lines[0].startswith("#"):
    dataset = HostList(name=lines[0].strip("#").strip())
    host_datasets.append(dataset)
    lines = lines[1:]
else:
    dataset = HostList(name="Default")
    host_datasets.append(dataset)

for line in lines:
    if line.startswith("#"):
        dataset = HostList(name=line.strip("#").strip())
        host_datasets.append(dataset)
    else:
        host = line.strip()
        if host and (host not in all_hosts):
            all_hosts.append(host)
            host_datasets[-1].hosts.append(host)

with open("providers.txt", "w", encoding="utf-8") as f:
    for dataset in host_datasets:
        f.write(f"# {dataset.name}\n")
        f.write("\n")
        dataset.hosts.sort()
        write_lines = "\n".join(dataset.hosts)
        f.write(write_lines)
        f.write("\n")
