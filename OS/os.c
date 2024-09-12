#include<iostream>
using namespace std;

struct Process {
    int pid;
    int at; // arrival time
    int bt; // burst time
    int wt; // waiting time
    int tat; // turn around time
    int ct; // completion time
};

void findWaitingTime(Process processes[], int n) {
    processes[0].wt = 0;
    for (int i = 1; i < n; i++) {
        processes[i].wt = processes[i-1].ct - processes[i].at;
        if (processes[i].wt < 0) {
            processes[i].wt = 0;
        }
    }
}

void findTurnAroundTime(Process processes[], int n) {
    for (int i = 0; i < n; i++) {
        processes[i].tat = processes[i].bt + processes[i].wt;
    }
}

void findavgTime(Process processes[], int n) {
    float total_wt = 0;
    float total_tat = 0;
    for (int i = 0; i < n; i++) {
        total_wt += processes[i].wt;
        total_tat += processes[i].tat;
    }
    float avg_wt = total_wt / n;
    float avg_tat = total_tat / n;
    cout << "Average waiting time = " << avg_wt << endl;
    cout << "Average turn around time = " << avg_tat << endl;
}

int main() {
    int n;
    cout << "Enter the number of processes: ";
    cin >> n;

    Process processes[n];
    for (int i = 0; i < n; i++) {
        cout << "Enter arrival time and burst time for process " << i+1 << ": ";
        cin >> processes[i].at >> processes[i].bt;
        processes[i].pid = i + 1;
    }

    // Sort processes based on arrival time
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            if (processes[i].at > processes[j].at) {
                swap(processes[i], processes[j]);
            }
        }
    }

    // Calculate completion time
    processes[0].ct = processes[0].at + processes[0].bt;
    for (int i = 1; i < n; i++) {
        processes[i].ct = processes[i-1].ct + processes[i].bt;
    }

    findWaitingTime(processes, n);
    findTurnAroundTime(processes, n);
    findavgTime(processes, n);

    return 0;
}
