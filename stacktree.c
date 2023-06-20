#include <stdio.h>
#include <stdlib.h>

// Structure representing a node in the binary tree
struct Node {
    int data;
    struct Node* left;
    struct Node* right;
};

// Function to create a new node
struct Node* createNode(int data) {
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->left = NULL;
    newNode->right = NULL;
    return newNode;
}

// Function to perform postorder traversal of the binary tree using a stack
void postorderTraversal(struct Node* root) {
    if (root == NULL)
        return;

    struct Node* stack1[100];
    struct Node* stack2[100];
    int top1 = -1;
    int top2 = -1;

    stack1[++top1] = root;

    while (top1 >= 0) {
        struct Node* current = stack1[top1--];
        stack2[++top2] = current;

        if (current->left)
            stack1[++top1] = current->left;

        if (current->right)
            stack1[++top1] = current->right;
    }

    while (top2 >= 0) {
        struct Node* current = stack2[top2--];
        printf("%d ", current->data);
    }
}

int main() {
    struct Node* root = createNode(1);
    root->left = createNode(2);
    root->right = createNode(3);
    root->left->left = createNode(4);
    root->left->right = createNode(5);

    printf("Postorder Traversal: ");
    postorderTraversal(root);
    printf("\n");

    return 0;
}
